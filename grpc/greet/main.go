package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/hwipl/snippets-go/grpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Greet(_ context.Context, g *pb.Greeting) (*pb.Greeting, error) {
	log.Println("Received request from client:", g)
	r := &pb.Greeting{
		FromName: "server",
		ToName:   g.FromName,
		Text:     g.Text,
	}
	log.Println("Sending reply to client:", r)
	return r, nil
}

func main() {
	// server
	// start tcp listenrer
	listener, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// start serving grpc
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Server listening on %v", listener.Addr())
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// client
	// connect to server
	conn, err := grpc.NewClient(listener.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// send request to server and get reply
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := &pb.Greeting{
		FromName: "client",
		ToName:   "server",
		Text:     "hi",
	}
	log.Println("Sending request to server:", request)
	reply, err := c.Greet(ctx, request)
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Println("Received reply from server:", reply)
}
