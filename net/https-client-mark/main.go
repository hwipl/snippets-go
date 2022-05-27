package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	// create dialer
	mark := 42

	// set socket option function
	var soerr error
	setsockopt := func(fd uintptr) {
		log.Println("Setting SO_MARK to", mark)
		soerr = unix.SetsockoptInt(
			int(fd),
			unix.SOL_SOCKET,
			unix.SO_MARK,
			mark,
		)
		if soerr != nil {
			log.Println("error setting SO_MARK:", soerr)
		}
	}

	// control function that sets socket option on raw connection
	control := func(network, address string, c syscall.RawConn) error {
		if err := c.Control(setsockopt); err != nil {
			return err
		}
		return soerr
	}

	// create dialer
	dialer := &net.Dialer{
		Control: control,
	}

	// create http client
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}

	// run http get request
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		log.Fatal("error during get: ", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Fatal(err)
	}
	log.Println("GET successful")
}
