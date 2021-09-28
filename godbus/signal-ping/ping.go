package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

const (
	// object path, interface, ping and pong members
	path  = "/org/ping/Ping"
	iface = "org.ping.Ping"
	ping  = iface + ".Ping"
	pong  = iface + ".Pong"
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
		dbus.WithMatchObjectPath(path),
		dbus.WithMatchInterface(iface),
	); err != nil {
		log.Fatal(err)
	}

	// create channel for signals
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	// send ping request
	conn.Emit(path, ping, "PING")

	// handle incoming signals
	name := conn.Names()[0]
	for s := range c {
		if s.Sender == name {
			// filter own signals
			continue
		}

		switch s.Name {
		case ping:
			// incoming ping request
			log.Println("PING from", s.Sender)
			conn.Emit(path, pong, "PONG")
		case pong:
			// incoming ping reply
			log.Println("PONG from", s.Sender)
		}
	}
}
