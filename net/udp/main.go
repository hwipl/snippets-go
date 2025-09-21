package main

import (
	"log"
	"net"
)

// runServer runs the udp server.
func runServer(listen net.PacketConn) {
	defer func() { _ = listen.Close() }()

	for {
		b := make([]byte, 2048)
		n, addr, err := listen.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Server handling new message from:", addr)
		listen.WriteTo(b[:n], addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// runClient runs the udp client.
func runClient(addr string) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	log.Println("Client sending message to server")
	if _, err := conn.Write([]byte("hi")); err != nil {
		log.Fatal(err)
	}

	log.Println("Client receiving reply from server")
	b := make([]byte, 2048)
	n, err := conn.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client received reply from server:", string(b[:n]))
}

func main() {
	// start server
	listen, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on", listen.LocalAddr())
	go runServer(listen)

	// start client
	runClient(listen.LocalAddr().String())
}
