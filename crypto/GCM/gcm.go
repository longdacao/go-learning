package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return cipherText, nil
}

func decrypt(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < 12 {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := cipherText[:12]
	cipherText = cipherText[12:]

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func main() {
	key := []byte("0123456789ABCDEF0123456789ABCDEF") // 32-byte key for AES-256
	plainText := []byte("Hello, GCM!")

	fmt.Printf("Original Text: %s\n", plainText)

	cipherText, err := encrypt(plainText, key)
	if err != nil {
		fmt.Printf("Encryption error: %v\n", err)
		return
	}

	fmt.Printf("Encrypted Text: %x\n", cipherText)

	decryptedText, err := decrypt(cipherText, key)
	if err != nil {
		fmt.Printf("Decryption error: %v\n", err)
		return
	}

	fmt.Printf("Decrypted Text: %s\n", decryptedText)
}
