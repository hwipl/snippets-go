package main

import "flag"

func main() {
	addr := flag.String("addr", ":8080", "set server `address`")
	name := flag.String("name", "", "set `name` of client")
	flag.Parse()

	if *name != "" {
		runClient(*addr, *name)
		return
	}

	runServer(*addr)
}
