package main

import (
	"flag"
	"log"
	"strings"

	"github.com/miekg/dns"
)

var (
	// default local server address, protocol (udp/tcp), remote server
	address  = "127.0.0.1:5353"
	protocol = "udp"
	remote   = "8.8.8.8:53"
)

// parseCommandLine parses the command line arguments
func parseCommandLine() {
	flag.Parse()

	// parse local server address
	a := flag.Arg(0)
	if a != "" {
		address = strings.ToLower(a)
	}

	// parse protocol
	p := flag.Arg(1)
	if p != "" {
		protocol = strings.ToLower(p)
	}

	// parse remote server address
	r := flag.Arg(2)
	if p != "" {
		remote = strings.ToLower(r)
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	// log request
	for _, q := range r.Question {
		log.Println("Question:", q)
	}

	// forward request to remote server
	r, err := dns.Exchange(r, remote)
	if err != nil {
		log.Fatal(err)
	}

	// log reply
	for _, a := range r.Answer {
		log.Println("Answer:", a)
	}

	// send reply back to client
	if err := w.WriteMsg(r); err != nil {
		log.Println(err)
	}
}

func main() {
	// parse command line arguments
	parseCommandLine()

	// register handler and start dns server
	dns.HandleFunc(".", handleRequest)
	err := dns.ListenAndServe(address, protocol, nil)
	if err != nil {
		log.Fatal(err)
	}
}
