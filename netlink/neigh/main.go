package main

import (
	"fmt"
	"log"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

func main() {
	// create channel for neigh updates and done, subscribe to neigh updates
	updates := make(chan netlink.NeighUpdate)
	done := make(chan struct{})
	err := netlink.NeighSubscribe(updates, done)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		done <- struct{}{}
	}()

	// handle neigh updates, print neighbor for the following event types:
	// - new neighbor
	// - delete neighbor
	for u := range updates {
		neigh := fmt.Sprintf("ip: %s, mac: %s, ifindex: %d", u.IP,
			u.HardwareAddr, u.LinkIndex)
		switch u.Type {
		case unix.RTM_NEWNEIGH:
			log.Println("NEW NEIGH:", neigh)
		case unix.RTM_DELNEIGH:
			log.Println("DEL NEIGH:", neigh)
		}
	}
}
