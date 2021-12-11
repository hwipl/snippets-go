package main

import (
	"log"
	"os"
	"strings"

	"github.com/asaskevich/govalidator"
)

func main() {
	// get hash from command line
	if len(os.Args) < 2 {
		log.Fatal("no hash specified")
	}
	hash := os.Args[1]

	// remove unwanted characters from hash
	hash = strings.TrimSpace(hash)
	hash = strings.ReplaceAll(hash, ":", "")
	hash = strings.ReplaceAll(hash, "-", "")

	// check hash type
	algs := []string{"md4", "md5", "sha1", "sha256", "sha3842", "sha512",
		"ripemd128", "ripemd160", "tiger128", "tiger160", "tiger192",
		"crc32", "crc32b"}
	for _, alg := range algs {
		if govalidator.IsHash(hash, alg) {
			log.Println(hash, "is", alg)
		}
	}
}
