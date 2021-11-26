package main

import (
	"log"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

func main() {
	// create channel for route updates and done, subscribe to route updates
	updates := make(chan netlink.RouteUpdate)
	done := make(chan struct{})
	err := netlink.RouteSubscribe(updates, done)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		done <- struct{}{}
	}()

	// handle route updates, print route for the following event types:
	// - new route
	// - delete route
	for u := range updates {
		switch u.Type {
		case unix.RTM_NEWROUTE:
			log.Println("NEW ROUTE:", u.Route)
		case unix.RTM_DELROUTE:
			log.Println("DEL ROUTE:", u.Route)
		}
	}
}
