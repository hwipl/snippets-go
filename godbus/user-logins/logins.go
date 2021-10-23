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

// parseUserSignal parses a UserNew or UserRemoved signal and returns the uid
// and object path
func parseUserSignal(s *dbus.Signal) (uint32, dbus.ObjectPath) {
	uid, ok := s.Body[0].(uint32)
	if !ok {
		log.Fatal("error parsing uid in user signal")
	}
	path, ok := s.Body[1].(dbus.ObjectPath)
	if !ok {
		log.Fatal("error parsing object path user signal")
	}
	return uid, path
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
		log.Printf("Currently logged in user: %d (name: \"%s\")",
			user.UID, user.Name)
	}

	// handle login signals
	for s := range c {
		switch s.Name {
		case userNew:
			// handle user new signal
			uid, _ := parseUserSignal(s)
			log.Printf("User login: %d", uid)
		case userRemoved:
			// handle user removed signal
			uid, _ := parseUserSignal(s)
			log.Printf("User logout: %d", uid)
		}
	}
}
