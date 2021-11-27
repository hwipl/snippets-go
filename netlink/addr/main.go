package main

import (
	"fmt"
	"log"

	"github.com/vishvananda/netlink"
)

func main() {
	// create channel for addr updates and done, subscribe to addr updates
	updates := make(chan netlink.AddrUpdate)
	done := make(chan struct{})
	err := netlink.AddrSubscribe(updates, done)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		done <- struct{}{}
	}()

	// handle addr updates, print address and interface index
	for u := range updates {
		addr := fmt.Sprintf("%s (ifindex: %d)", &u.LinkAddress,
			u.LinkIndex)
		if u.NewAddr {
			log.Println("New Addr:", addr)
		} else {
			log.Println("Del Addr:", addr)
		}
	}
}
