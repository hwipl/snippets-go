package main

import (
	"context"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// run a simple ls command to completion and read output
	lsCmd := exec.Command("ls")
	lsOut, err := lsCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range strings.Split(string(lsOut), "\n") {
		if s == "" {
			continue
		}
		log.Println(s)
	}

	// run a command with periodic output and read output
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)
	defer cancel()

	watchCmd := exec.CommandContext(ctx, "watch", "ls")
	stdout, err := watchCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := watchCmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Println("Running command for 10 seconds")
	if _, err := io.Copy(io.Discard, stdout); err != nil {
		log.Fatal(err)
	}
	if err := watchCmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
