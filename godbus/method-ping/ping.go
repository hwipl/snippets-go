package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

const (
	// object path, interface, ping member
	path       = "/org/ping/Ping"
	iface      = "org.ping.Ping"
	pingMethod = iface + ".Ping"
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
	conn.Export(ping{}, path, iface)

	// request name
	reply, err := conn.RequestName(iface, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		// name already taken, send ping request to it
		s := ""
		err := conn.Object(iface, path).Call(pingMethod, 0).Store(&s)
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
