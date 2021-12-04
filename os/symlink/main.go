package main

import (
	"log"
	"os"
)

func main() {
	// create a symlink
	err := os.Symlink("/tmp", "test")
	if err != nil {
		log.Println(err)
	}

	// read symlink
	link, err := os.Readlink("test")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("test ->", link)
}
