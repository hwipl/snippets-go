package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Greeting struct {
	Greeting string
	From     string
	To       []string
}

func main() {
	hw1 := &Greeting{
		Greeting: "Hello",
		From:     "yaml",
		To:       []string{"world", "*"},
	}
	b, err := yaml.Marshal(hw1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	hw2 := &Greeting{}
	err = yaml.Unmarshal(b, hw2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", hw2)
}
