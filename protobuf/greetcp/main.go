package main

// compiling protocol buffers:
// protoc --go_out=. --go_opt=paths=source_relative greetpb/greet.proto

import (
	"encoding/binary"
	"io"
	"log"
	"net"

	pb "github.com/hwipl/snippets-go/protobuf/greetcp/greetpb"
	"google.golang.org/protobuf/proto"
)

// handleClient handles a client connection to the server.
func handleClient(conn net.Conn) {
	defer func() { _ = conn.Close() }()

	log.Println("Server handling new client connection:", conn.RemoteAddr())

	// read message length
	l := make([]byte, 4)
	if _, err := io.ReadFull(conn, l); err != nil {
		log.Fatal(err)
	}

	// read message
	b := make([]byte, binary.LittleEndian.Uint32(l))
	if _, err := io.ReadFull(conn, b); err != nil {
		log.Fatal(err)
	}

	// parse message
	g := &pb.Greeting{}
	if err := proto.Unmarshal(b, g); err != nil {
		log.Fatal(err)
	}
	log.Println("Server got message from client:", g)

	// create reply
	g.ToName = g.FromName
	g.FromName = "server"

	b, err := proto.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}

	// send reply length
	binary.LittleEndian.PutUint32(l, uint32(len(b)))
	if _, err := conn.Write(l); err != nil {
		log.Fatal(err)
	}

	// send reply
	if _, err := conn.Write(b); err != nil {
		log.Fatal(err)
	}
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

	// create message
	g1 := &pb.Greeting{
		FromName: "client",
		ToName:   "server",
		Text:     "hi",
	}
	b, err := proto.Marshal(g1)
	if err != nil {
		log.Fatal(err)
	}

	// send message length
	l := make([]byte, 4)
	binary.LittleEndian.PutUint32(l, uint32(len(b)))
	if _, err := conn.Write(l); err != nil {
		log.Fatal(err)
	}

	// send message
	if _, err := conn.Write(b); err != nil {
		log.Fatal(err)
	}

	// read reply length
	if _, err := io.ReadFull(conn, l); err != nil {
		log.Fatal(err)
	}

	// read reply
	b = make([]byte, binary.LittleEndian.Uint32(l))
	if _, err = io.ReadFull(conn, b); err != nil {
		log.Fatal(err)
	}

	// parse reply
	g2 := &pb.Greeting{}
	if err := proto.Unmarshal(b, g2); err != nil {
		log.Fatal(err)
	}
	log.Println("Client got reply from server:", g2)
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
