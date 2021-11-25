package main

import (
	"fmt"
	"log"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

func main() {
	// create channel for link updates and done, subscribe to link updates
	updates := make(chan netlink.LinkUpdate)
	done := make(chan struct{})
	err := netlink.LinkSubscribe(updates, done)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		done <- struct{}{}
	}()

	// handle link updates, print device name and index for the following
	// event types:
	// - new link
	// - delete link
	for u := range updates {
		attrs := u.Link.Attrs()
		device := fmt.Sprintf("%s (index: %d)", attrs.Name, attrs.Index)
		switch u.Header.Type {
		case unix.RTM_NEWLINK:
			log.Println("NEW LINK:", device)
		case unix.RTM_DELLINK:
			log.Println("DEL LINK:", device)
		}
	}
}
