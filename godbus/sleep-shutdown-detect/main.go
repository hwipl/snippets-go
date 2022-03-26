// Detect suspend/hibernate and reboot/poweroff

// Based on:
// https://www.freedesktop.org/wiki/Software/systemd/logind/
// and
// $ gdbus introspect --system --dest org.freedesktop.login1 \
//     --object-path /org/freedesktop/login1
// and
// $ gdbus introspect --system --dest org.freedesktop.login1 \
//     --object-path /org/freedesktop/login1/user/_1000

package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

const (
	// object path, destination, interface, signals, methods, properties
	path               = "/org/freedesktop/login1"
	dest               = "org.freedesktop.login1"
	iface              = dest + ".Manager"
	prepareForSleep    = iface + ".PrepareForSleep"
	prepareForShutdown = iface + ".PrepareForShutdown"
)

func main() {
	// connect to system bus
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

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

	// handle login signals
	for s := range c {
		switch s.Name {
		case prepareForShutdown:
			// handle prepare for shutdown signal
			shutdown, ok := s.Body[0].(bool)
			if !ok {
				log.Fatal("error parsing prepare for shutdown signal")
			}
			if shutdown {
				log.Println("System going into shutdown")
			} else {
				log.Println("System resuming from shutdown")
			}
		case prepareForSleep:
			// handle prepare for sleep signal
			sleep, ok := s.Body[0].(bool)
			if !ok {
				log.Fatal("error parsing prepare for sleep signal")
			}
			if sleep {
				log.Println("System going into sleep")
			} else {
				log.Println("System resuming from sleep")
			}
		}
	}
}
