package main

import (
	"log"
	"os"

	"github.com/asaskevich/govalidator"
)

// printType prints addr and its type
func printType(addr, typ string) {
	log.Println(addr, "is a", typ)
}

func main() {
	// get address from command line
	if len(os.Args) < 2 {
		log.Fatal("no address specified")
	}
	addr := os.Args[1]

	// check address type
	if govalidator.IsURL(addr) {
		printType(addr, "URL")
	}
	if govalidator.IsHost(addr) {
		printType(addr, "hostname")
	}
	if govalidator.IsDNSName(addr) {
		printType(addr, "domain name")
	}
	if govalidator.IsPort(addr) {
		printType(addr, "port number")
	}
	if govalidator.IsIP(addr) {
		printType(addr, "IP address")
	}
	if govalidator.IsIPv4(addr) {
		printType(addr, "IPv4 address")
	}
	if govalidator.IsIPv6(addr) {
		printType(addr, "IPv6 address")
	}
	if govalidator.IsMAC(addr) {
		printType(addr, "MAC address")
	}
}
