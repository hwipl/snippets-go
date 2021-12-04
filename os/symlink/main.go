package main

import (
	"log"
	"os"
	"path/filepath"
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

	// read and evaluate symlink
	linkEval, err := filepath.EvalSymlinks("test")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("test ->", linkEval)
}
