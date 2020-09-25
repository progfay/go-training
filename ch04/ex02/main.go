package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	hashType := flag.String("type", "sha256", "type name of hash function (sha256, sha384, sha512)")
	flag.Parse()

	switch *hashType {
	case "sha256", "sha384", "sha512":
	default:
		panic(fmt.Errorf("invalid type: %s", *hashType))
	}

	fmt.Printf("type: %s\n", *hashType)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		switch *hashType {
		case "sha256":
			hash := sha256.Sum256(bytes)
			fmt.Println(hex.EncodeToString(hash[:]))

		case "sha384":
			hash := sha512.Sum384(bytes)
			fmt.Println(hex.EncodeToString(hash[:]))

		case "sha512":
			hash := sha512.Sum512(bytes)
			fmt.Println(hex.EncodeToString(hash[:]))
		}
	}
}
