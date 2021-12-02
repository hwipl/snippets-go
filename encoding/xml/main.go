package main

import (
	"encoding/xml"
	"log"
	"time"
)

// Data defines some data to be encoded as xml
type Data struct {
	Name  string
	Value string
	Count uint32
	Time  time.Time
}

func main() {
	// create data to be encoded as xml
	encode := Data{
		"hi",
		"hello world",
		23,
		time.Now().Local(),
	}
	log.Println("encode:", encode)

	// encode data as xml
	x, err := xml.Marshal(encode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("xml:", string(x))

	// decode data from xml
	var decode Data
	if err := xml.Unmarshal(x, &decode); err != nil {
		log.Println(err)
	}
	log.Println("decode:", decode)

	// check if encoded and decoded data are equal
	if encode == decode {
		log.Println("encode == decode")
	}
}
