// List currently logged in users and show new logins and logouts.
// Based on:
// https://www.freedesktop.org/wiki/Software/systemd/logind/
// and
// $ gdbus introspect --system --dest org.freedesktop.login1 \
//     --object-path /org/freedesktop/login1

package main

import (
	"log"

	"github.com/godbus/dbus/v5"
)

const (
	// object path, destination, interface, signals, methods
	path        = "/org/freedesktop/login1"
	dest        = "org.freedesktop.login1"
	iface       = dest + ".Manager"
	userNew     = iface + ".UserNew"
	userRemoved = iface + ".UserRemoved"
	listUsers   = iface + ".ListUsers"
)

// User is a currently logged in user
type User struct {
	UID  uint32
	Name string
	Path dbus.ObjectPath
}

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

	// request currently logged in users
	var users []User
	err = conn.Object(dest, path).Call(listUsers, 0).Store(&users)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		log.Printf("Currently logged in user: %s (uid: %d)", user.Name,
			user.UID)
	}

	// handle login signals
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
