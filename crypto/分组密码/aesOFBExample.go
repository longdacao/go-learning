package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte("0123456789abcdef")
	iv := []byte("1234567890abcdef")

	// 加密
	plaintext := []byte("Hello, OFB mode!")
	ciphertext := make([]byte, len(plaintext))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	fmt.Printf("Plaintext:  %x\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// 解密
	decrypted := make([]byte, len(ciphertext))

	stream = cipher.NewOFB(block, iv)
	stream.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Decrypted:  %s\n", decrypted)
}
