package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	hashWidth = flag.Int("width", 256, "hash width (384 or 512. Default is 256)")
)

func main() {
	// parse the flag
	flag.Parse()

	// get a string from user whose hash you want to find
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the string: \n")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading the string: %v", err)
		os.Exit(1)
	}

	var hashFunction func(b []byte) []byte

	switch *hashWidth {
	case 384:
		hashFunction = func(b []byte) []byte {
			res := sha512.Sum384(b)
			return res[:]
		}

	case 512:
		hashFunction = func(b []byte) []byte {
			res := sha512.Sum512(b)
			return res[:]
		}

	default:
		hashFunction = func(b []byte) []byte {
			res := sha256.Sum256(b)
			return res[:]
		}
	}

	fmt.Printf("%x\n", hashFunction([]byte(str)))
}
