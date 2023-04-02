package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("0123456789abcdef")
	plaintext := []byte("exampleplaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(plaintext))
	block.Encrypt(ciphertext, plaintext)

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted := make([]byte, len(ciphertext))
	block.Decrypt(decrypted, ciphertext)

	fmt.Printf("Decrypted: %s\n", decrypted)
}
