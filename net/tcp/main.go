package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

// handleClient handles a client connection to the server.
func handleClient(conn net.Conn) {
	defer func() { _ = conn.Close() }()

	log.Println("Server handling new client connection:", conn.RemoteAddr())
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatal(err)
	}
	log.Println("Server removing client connection:", conn.RemoteAddr())

}

// runServer runs the tcp server.
func runServer(listen net.Listener) {
	defer func() { _ = listen.Close() }()

	log.Println("Server listening on", listen.Addr())
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleClient(conn)
	}
}

// runClient runs the tcp client.
func runClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	log.Println("Client connected to server:", conn.RemoteAddr())
	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)

	log.Println("Client sending message to server")
	w.WriteString("hi\n")
	w.Flush()

	log.Println("Client receiving reply from server")
	reply, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Client received reply from server:", reply)
}

func main() {
	// start server
	listen, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	go runServer(listen)

	// start client
	runClient(listen.Addr().String())
}
