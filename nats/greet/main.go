package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// message is a greeting message.
type message struct {
	From string
	To   string
	Text string
}

func main() {
	// command line arguments
	name := flag.String("name", "client", "set name")
	flag.Parse()

	// connect to server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// subscribe to topic
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("hi", ch)
	defer sub.Unsubscribe()

	// timer
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-ch:
			// parse and show message
			m := &message{}
			if err := json.Unmarshal(msg.Data, m); err != nil {
				log.Fatal(err)
			}
			log.Printf("%s: %s", m.From, m.Text)

			// reply with a greeting
			if m.From != *name && m.To == "" {
				reply := &message{
					From: *name,
					To:   m.From,
					Text: fmt.Sprintf("hi, %s", m.From),
				}
				b, err := json.Marshal(reply)
				if err != nil {
					log.Fatal(err)
				}
				if err := nc.Publish("hi", b); err != nil {
					log.Fatal(err)
				}
			}
		case <-ticker.C:
			// send greeting
			m := &message{
				From: *name,
				Text: "hi",
			}
			b, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}
			if err := nc.Publish("hi", b); err != nil {
				log.Fatal(err)
			}
		}
	}
}
