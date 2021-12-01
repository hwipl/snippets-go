package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
)

// handleClient handles a client connection to the server
func handleClient(conn net.Conn) {
	log.Println("Handle new client connection:", conn.RemoteAddr())
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatal(err)
	}
	log.Println("Remove client connection:", conn.RemoteAddr())

}

// runServer runs a unix socket server
func runServer() {
	listen, err := net.Listen("unix", "unix.sock")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleClient(conn)
	}
}

// runClient runs a unix socket client
func runClient() {
	conn, err := net.Dial("unix", "unix.sock")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to server:", conn.RemoteAddr())
	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)

	log.Println("Sending message to server")
	w.WriteString("hi\n")
	w.Flush()

	log.Println("Receiving reply from server")
	reply, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Received reply from server:", reply)
}

func main() {
	// define and parse command line arguments
	server := flag.Bool("s", false, "run as server")
	flag.Parse()

	// run in server or client mode
	if *server {
		runServer()
		return
	}
	runClient()
}
