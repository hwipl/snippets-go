package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"log"
	"net/http"
)

func main() {
	// run http get request
	r, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	// get tls version from response
	if r.TLS == nil {
		log.Fatal("no tls connection state")
	}
	version := ""
	switch r.TLS.Version {
	case tls.VersionTLS10:
		version = "TLS 1.0"
	case tls.VersionTLS11:
		version = "TLS 1.1"
	case tls.VersionTLS12:
		version = "TLS 1.2"
	case tls.VersionTLS13:
		version = "TLS 1.3"
	default:
		log.Fatal("unsupported tls version")
	}

	// get fingerprint of server certificate
	cert := r.TLS.PeerCertificates[0]
	hash := sha256.Sum256(cert.Raw)
	fp := hex.EncodeToString(hash[:])

	// print server info
	log.Println(r.Status, version, r.TLS.ServerName, cert.Subject, fp)
}
