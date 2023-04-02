package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	var randomBytes [8]byte
	_, err := rand.Read(randomBytes[:])
	if err != nil {
		panic(err)
	}

	randomNumber := uint64(randomBytes[0]) |
		uint64(randomBytes[1])<<8 |
		uint64(randomBytes[2])<<16 |
		uint64(randomBytes[3])<<24 |
		uint64(randomBytes[4])<<32 |
		uint64(randomBytes[5])<<40 |
		uint64(randomBytes[6])<<48 |
		uint64(randomBytes[7])<<56

	fmt.Println(randomNumber)
}
