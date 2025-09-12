package main

// compiling protocol buffers:
// protoc --go_out=. --go_opt=paths=source_relative greetpb/greet.proto

import (
	"fmt"

	pb "github.com/hwipl/snippets-go/protobuf/greet/greetpb"
	"google.golang.org/protobuf/proto"
)

func main() {

	g1 := &pb.Greeting{
		FromName: "me",
		ToName:   "you",
		Text:     "hi",
	}
	b, err := proto.Marshal(g1)
	if err != nil {
		fmt.Println(err)
		return
	}

	g2 := &pb.Greeting{}
	if err := proto.Unmarshal(b, g2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(g2)
}
