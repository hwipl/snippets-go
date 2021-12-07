package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// handler sends greeting back to the client
func handler(greeting string, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		io.WriteString(w, fmt.Sprintf("%s, %s!\n", greeting, name))
	} else {
		io.WriteString(w, fmt.Sprintf("%s!\n", greeting))
	}
}

// hiHandler sends "hello" back to the client
func hiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hi")
	handler("hello", w, r)
}

// byeHandler sends "bye" back to the client
func byeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("bye")
	handler("bye", w, r)
}

func main() {
	// add handlers
	http.HandleFunc("/hi", hiHandler)
	http.HandleFunc("/bye", byeHandler)

	// start http server
	log.Fatal(http.ListenAndServe(":8080", nil))

}
