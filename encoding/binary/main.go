package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type header struct {
	Type   uint16
	Length uint32
}

func main() {
	buf := &bytes.Buffer{}

	// write header and data to buffer
	d := []byte("hello world")
	h := &header{
		Type:   1,
		Length: uint32(len(d)),
	}
	if err := binary.Write(buf, binary.LittleEndian, h); err != nil {
		log.Fatal(err)
	}
	if err := binary.Write(buf, binary.LittleEndian, d); err != nil {
		log.Fatal(err)
	}

	// read header and data from buffer
	h = &header{}
	if err := binary.Read(buf, binary.LittleEndian, h); err != nil {
		log.Fatal(err)
	}
	d = make([]byte, h.Length)
	//if err := binary.Read(buf, binary.LittleEndian, d); err != nil {
	if _, err := io.ReadFull(buf, d); err != nil {
		log.Fatal(err)
	}
	log.Printf("Type: %d, Length: %d, Data: %s\n", h.Type, h.Length, d)
}
