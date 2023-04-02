package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

const key = "01234567"

func main() {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	plaintext := []byte("Hello, world!")
	fmt.Printf("Plaintext: %s\n", plaintext)

	ciphertext := make([]byte, len(plaintext))
	feistelEncrypt(block, plaintext, ciphertext)

	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted := make([]byte, len(ciphertext))
	feistelDecrypt(block, ciphertext, decrypted)

	fmt.Printf("Decrypted: %s\n", decrypted)
}

func feistelEncrypt(block cipher.Block, plaintext, ciphertext []byte) {
	left := make([]byte, 4)
	right := make([]byte, 4)

	copy(left, plaintext[:4])
	copy(right, plaintext[4:])

	for i := 0; i < 16; i++ {
		tmp := make([]byte, 4)
		block.Encrypt(tmp, left)
		for j := 0; j < 4; j++ {
			tmp[j] ^= right[j]
		}
		copy(left, right)
		copy(right, tmp)
	}

	copy(ciphertext[:4], right)
	copy(ciphertext[4:], left)
}

func feistelDecrypt(block cipher.Block, ciphertext, decrypted []byte) {
	left := make([]byte, 4)
	right := make([]byte, 4)

	copy(left, ciphertext[4:])
	copy(right, ciphertext[:4])

	for i := 0; i < 16; i++ {
		tmp := make([]byte, 4)
		block.Decrypt(tmp, right)
		for j := 0; j < 4; j++ {
			tmp[j] ^= left[j]
		}
		copy(right, left)
		copy(left, tmp)
	}

	copy(decrypted, left)
	copy(decrypted[4:], right)
}
