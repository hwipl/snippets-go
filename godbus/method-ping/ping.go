package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

const (
	// object path
	path = "/org/ping/Ping"
)

// define ping interface methods
type ping struct{}

func (p ping) Ping(sender dbus.Sender) (string, *dbus.Error) {
	log.Println("PING from", sender)
	return "PONG", nil
}

func main() {
	// connect to session bus
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// export ping interface methods
	conn.Export(ping{}, path, "org.ping.Ping")

	// request name
	reply, err := conn.RequestName("org.ping.Ping",
		dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		// name already taken, send ping request to it
		s := ""
		err := conn.Object("org.ping.Ping",
			path).Call("org.ping.Ping.Ping", 0).Store(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(s)
		return

	}

	// got the name, handle ping requests
	log.Println("Waiting for ping requests...")
	select {}
}
