package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
)

var (
	// default server address and protocol (udp/tcp)
	address  = "127.0.0.1:5353"
	protocol = "udp"

	// default server replies
	defaultA    = net.ParseIP("127.0.0.1")
	defaultAAAA = net.ParseIP("::1")
)

// parseCommandLine parses the command line arguments
func parseCommandLine() {
	flag.Parse()

	// parse address
	a := flag.Arg(0)
	if a != "" {
		address = strings.ToLower(a)
	}

	// parse protocol
	p := flag.Arg(1)
	if p != "" {
		protocol = strings.ToLower(p)
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	log.Println(r)

	// create reply message
	msg := &dns.Msg{}
	msg.SetReply(r)

	// add requested record to reply message
	switch r.Question[0].Qtype {
	case dns.TypeA:
		// prepare A record
		rrA := &dns.A{
			Hdr: dns.RR_Header{
				Name:   r.Question[0].Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			A: defaultA,
		}

		// add it to answers
		msg.Answer = append(msg.Answer, rrA)
	case dns.TypeAAAA:
		// prepare AAAA record
		rrAAAA := &dns.AAAA{
			Hdr: dns.RR_Header{
				Name:   r.Question[0].Name,
				Rrtype: dns.TypeAAAA,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			AAAA: defaultAAAA,
		}

		// add it to answers
		msg.Answer = append(msg.Answer, rrAAAA)
	}

	// send reply back to client
	err := w.WriteMsg(msg)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// parse command line arguments
	parseCommandLine()

	// register handler and start dns server
	dns.HandleFunc("test.", handleRequest)
	err := dns.ListenAndServe(address, protocol, nil)
	if err != nil {
		log.Fatal(err)
	}
}
