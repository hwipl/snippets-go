package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// runClient runs as a client.
func runClient(addr, name string) {
	// handle interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// connect to server
	u := url.URL{Scheme: "ws", Host: addr, Path: "/greet"}
	log.Printf("Connecting to server: %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Could not connect to server:", err)
	}
	defer c.Close()

	// read messages from connection in other goroutine
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			// read message
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Could not read message:", err)
				return
			}
			m := string(message)
			log.Println(m)

			// send reply
			if strings.HasPrefix(m, "Hi, I'm ") {
				name := m[8:]
				m := fmt.Sprintf("Hi, %s", name)
				log.Println(m)

				if err := c.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
					log.Println("Could not write message:", err)
					return
				}
			}
		}
	}()

	// timer for sending messages
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			// goroutine for reading messages stopped, stop
			return
		case <-ticker.C:
			// timer, send message
			m := fmt.Sprintf("Hi, I'm %s", name)
			log.Println(m)
			if err := c.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
				log.Println("Could not write message:", err)
				return
			}
		case <-interrupt:
			// got interrupt signal, stop
			log.Println("Got interrupt signal, stopping...")

			// send close message and wait for server to close connection
			if err := c.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure,
					"")); err != nil {
				log.Println("Could not write close message:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
