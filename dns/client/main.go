package main

import (
	"flag"
	"log"
	"strings"

	"github.com/miekg/dns"
)

var (
	// default name server, query type and name
	server = "8.8.8.8:53"
	query  = "a"
	name   = "www.google.com"
)

// parseCommandLine parses the command line arguments
func parseCommandLine() {
	flag.Parse()

	// parse name
	n := flag.Arg(0)
	if n != "" {
		name = strings.ToLower(n)
	}

	// parse query type
	q := flag.Arg(1)
	if q != "" {
		query = strings.ToLower(q)
	}

	// parse server
	s := flag.Arg(2)
	if s != "" {
		server = strings.ToLower(s)
	}
}

func main() {
	// parse command line arguments
	parseCommandLine()

	// set query type
	typ := dns.TypeNone
	switch query {
	case "a":
		typ = dns.TypeA
	case "aaaa":
		typ = dns.TypeAAAA
	case "mx":
		typ = dns.TypeMX
	default:
		log.Fatal("unsupported query type: ", query)
	}

	// create and send message
	msg := &dns.Msg{}
	msg.SetQuestion(dns.Fqdn(name), typ)
	r, err := dns.Exchange(msg, server)
	if err != nil {
		log.Fatal(err)
	}

	// handle reply
	for _, a := range r.Answer {
		log.Println(a)
	}
}
