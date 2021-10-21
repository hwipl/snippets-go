package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

const (
	// object path, interface
	path        = "/org/freedesktop/login1"
	iface       = "org.freedesktop.login1.Manager"
	userNew     = iface + ".UserNew"
	userRemoved = iface + ".UserRemoved"
)

func main() {
	// connect to system bus
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// subscribe to login signals
	if err = conn.AddMatchSignal(
		dbus.WithMatchObjectPath(path),
		dbus.WithMatchInterface(iface),
	); err != nil {
		log.Fatal(err)
	}

	// create channel for signals
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	// handle login signals, based on:
	// https://www.freedesktop.org/wiki/Software/systemd/logind/
	// and
	// $ gdbus introspect --system --dest org.freedesktop.login1 \
	//     --object-path /org/freedesktop/login1
	for s := range c {
		switch s.Name {
		case userNew:
			// handle user new signal
			log.Println("User new with uid:", s.Body[0])
		case userRemoved:
			// handle user removed signal
			log.Println("User removed with uid:", s.Body[0])
		}
	}
}
