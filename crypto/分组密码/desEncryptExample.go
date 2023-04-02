package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte("qwertyui")
	plaintext := []byte("Hello World!!")

	// 创建DES加密算法的块
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 对明文进行PKCS5填充
	padding := 8 - len(plaintext)%8
	for i := 0; i < padding; i++ {
		plaintext = append(plaintext, byte(padding))
	}

	// 创建CBC模式的加密器
	iv := []byte("12345678")
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密明文
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
}
