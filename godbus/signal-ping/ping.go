package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

func main() {
	// connect to session bus
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// subscribe to specific signals
	if err = conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/org/ping/Ping"),
		dbus.WithMatchInterface("org.ping.Ping"),
	); err != nil {
		log.Fatal(err)
	}

	// create channel for signals
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	// send ping request
	conn.Emit("/org/ping/Ping", "org.ping.Ping.Ping", "PING")

	// handle incoming signals
	name := conn.Names()[0]
	for s := range c {
		if s.Sender == name {
			// filter own signals
			continue
		}

		switch s.Name {
		case "org.ping.Ping.Ping":
			// incoming ping request
			log.Println("PING from", s.Sender)
			conn.Emit("/org/ping/Ping", "org.ping.Ping.Pong",
				"PONG")
		case "org.ping.Ping.Pong":
			// incoming ping reply
			log.Println("PONG from", s.Sender)
		}
	}
}
