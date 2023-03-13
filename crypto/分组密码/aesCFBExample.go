package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte("0123456789abcdef")       // 16 字节的 key
	plaintext := []byte("Hello, CFB mode!") // 待加密的数据

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	fmt.Printf("Plaintext:  %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// 解密
	stream = cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	stream.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Decrypted:  %s\n", decrypted)
}
