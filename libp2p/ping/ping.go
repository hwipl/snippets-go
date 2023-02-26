// simple ping program that listens on the loopback interface on a random tcp
// port. Without command line arguments, the programs simply waits for incoming
// streams with the custom protocol id /hello/world/0.0.1, reads a message from
// it and writes it back to the stream. If first command line argument is
// given, it is treated as peer address and the program connects to it, sends a
// message, reads the reply and stops.

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// helloWorldHandler handles the hello world protocol
func helloWorldHandler(s network.Stream) {
	// read message
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		log.Println(err)
		s.Reset()
		return
	}
	log.Printf("message from %s: %s\n", s.ID(), str)

	// send message back
	_, err = s.Write([]byte(str))
	if err != nil {
		log.Println(err)
		s.Reset()
		return
	}

	// close stream
	if err := s.Close(); err != nil {
		log.Println(err)
	}

	return
}

// initNode initializes this node
func initNode() host.Host {
	// start a libp2p node that listens on a random local TCP port
	node, err :=
		libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		log.Fatal(err)
	}

	// configure our own ping protocol
	node.SetStreamHandler("/hello/world/0.0.1", helloWorldHandler)

	// print the node's PeerInfo in multiaddr format
	peerInfo := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	fmt.Println("libp2p node address:", addrs[0])

	return node
}

// listen simply waits for incoming streams until interrupted
func listen() {
	// initialize node
	node := initNode()

	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	// shut the node down
	if err := node.Close(); err != nil {
		log.Fatal(err)
	}
}

// connect connects to the peer identified by the multiaddr addr, sends a
// message and reads reply
func connect(addr multiaddr.Multiaddr) {
	// create a background context (i.e. one that never cancels)
	ctx := context.Background()

	// initialize node
	node := initNode()

	// parse peer multiaddress
	peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		log.Fatal(err)
	}

	// add peer to peer store
	node.Peerstore().AddAddr(peerInfo.ID, peerInfo.Addrs[0],
		peerstore.PermanentAddrTTL)

	// create new stream for hello world protocol
	s, err := node.NewStream(ctx, peerInfo.ID, "/hello/world/0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	// write message to stream
	_, err = s.Write([]byte("hi\n"))
	if err != nil {
		log.Println(err)
		s.Reset()
		return
	}

	// read reply
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		log.Println(err)
		s.Reset()
		return
	}
	log.Printf("reply from %s: %s\n", s.ID(), str)

	// close stream
	if err := s.Close(); err != nil {
		log.Fatal(err)
	}

	// shut down node
	if err := node.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// handle command line arguments
	if len(os.Args) > 1 {
		// first command line argument given, connect to peer in it
		addr, err := multiaddr.NewMultiaddr(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		connect(addr)
	} else {
		// no command line argument, just wait for incoming streams
		listen()
	}
}
