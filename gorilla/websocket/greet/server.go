package main

import (
	"log"
	"net/http"
	"slices"

	"github.com/gorilla/websocket"
)

// message is a greeting message.
type message struct {
	conn *websocket.Conn
	b    []byte
}

// clients manages all clients.
type clients struct {
	conns      []*websocket.Conn
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	messages   chan *message
}

// handle handles client connections and incoming messages.
func (c *clients) handle() {
	for {
		select {
		case conn := <-c.register:
			// add new client connection
			if slices.Contains(c.conns, conn) {
				break
			}
			log.Println("Registering client")
			c.conns = append(c.conns, conn)

		case conn := <-c.unregister:
			// remove existing client connection
			if i := slices.Index(c.conns, conn); i != -1 {
				log.Println("Unregistering client")
				c.conns = slices.Delete(c.conns, i, i+1)
			}
		case message := <-c.messages:
			// forward incoming message to other clients
			for _, conn := range c.conns {
				if conn == message.conn {
					continue
				}
				err := conn.WriteMessage(websocket.TextMessage, message.b)
				if err != nil {
					log.Println("Could not write message:", err)
					continue
				}
			}
		}
	}
}

var upgrader = websocket.Upgrader{}

// handleGreet handles the websocket connection to /greet.
func handleGreet(clients *clients, w http.ResponseWriter, r *http.Request) {
	// upgrade to ws
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Could not upgrade connection:", err)
		return
	}
	defer c.Close()

	// register client
	clients.register <- c
	defer func() {
		clients.unregister <- c
	}()

	// handle incoming messages
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Could not read message:", err)
			break
		}
		if mt != websocket.TextMessage {
			continue
		}
		log.Printf("Received message from client: %s", msg)
		m := &message{
			conn: c,
			b:    msg,
		}
		clients.messages <- m
	}
}

// runServer runs as a server.
func runServer(addr string) {
	cs := &clients{
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		messages:   make(chan *message),
	}
	go cs.handle()

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		handleGreet(cs, w, r)
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}
