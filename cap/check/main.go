package main

import (
	"log"

	"kernel.org/pub/linux/libs/security/libcap/cap"
)

func main() {
	// check permitted capabilities
	c := cap.GetProc()
	netAdmin, err := c.GetFlag(cap.Permitted, cap.NET_ADMIN)
	if err != nil {
		log.Fatal(err)
	}
	if !netAdmin {
		log.Fatal("Capability NET_ADMIN not set, " +
			"use \"setcap cap_net_admin=p\"")
	}

	// set effective capabilities
	if err := c.SetFlag(cap.Effective, true, cap.NET_ADMIN); err != nil {
		log.Fatal(err)
	}
	if err := c.SetProc(); err != nil {
		log.Fatal(err)
	}

	// drop capabilities
	empty := cap.NewSet()
	if err := empty.SetProc(); err != nil {
		log.Fatal(err)
	}

	// show differences
	n := cap.GetProc()
	log.Println(n.Cf(c))
}
