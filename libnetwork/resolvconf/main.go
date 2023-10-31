package main

import (
	"bytes"
	"log"

	"github.com/docker/docker/libnetwork/resolvconf"
)

func main() {
	// get default resolv.conf
	r, err := resolvconf.Get()
	if err != nil {
		log.Fatal(err)
	}

	// get nameservers
	for _, ns := range resolvconf.GetNameservers(r.Content, resolvconf.IP) {
		log.Println("Nameserver:", ns)
	}

	// get nameservers in CIDR notation
	for _, ns := range resolvconf.GetNameserversAsCIDR(r.Content) {
		log.Println("Nameserver (CIDR):", ns)
	}

	// get options
	for _, op := range resolvconf.GetOptions(r.Content) {
		log.Println("Option:", op)
	}

	// get search domains
	for _, sd := range resolvconf.GetSearchDomains(r.Content) {
		log.Println("Search Domain:", sd)
	}

	// get specific resolv.conf
	s, err := resolvconf.GetSpecific("/etc/resolv.conf")
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(s.Hash, r.Hash) {
		log.Println("resolv.confs are not equal")
	} else {
		log.Println("resolv.confs are equal")
	}
}
