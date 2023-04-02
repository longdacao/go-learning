package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	// Generate P-256 curve
	p256 := elliptic.P256()

	// Generate private key
	priv, err := ecdsa.GenerateKey(p256, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// Generate public key
	pub := priv.PublicKey

	// Generate random message to sign
	msg := []byte("Hello, world!")

	// Sign message
	r, s, err := ecdsa.Sign(rand.Reader, priv, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Verify signature
	if ecdsa.Verify(&pub, msg, r, s) {
		fmt.Println("Signature is valid.")
	} else {
		fmt.Println("Signature is invalid.")
	}
}
