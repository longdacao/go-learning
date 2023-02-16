package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte("12345678")
	plaintext := []byte("Hello World!!")

	// 加密明文
	ciphertext, err := desEncrypt(key, plaintext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// 解密密文
	plaintext2, err := desDecrypt(key, ciphertext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Plaintext2: %s\n", plaintext2)
}

// DES加密函数
func desEncrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	padding := 8 - len(plaintext)%8
	for i := 0; i < padding; i++ {
		plaintext = append(plaintext, byte(padding))
	}

	ciphertext := make([]byte, len(plaintext))
	// iv := ciphertext[:8]
	iv := []byte("12345678")

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// DES解密函数
func desDecrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// iv := ciphertext[:8]
	iv := []byte("12345678")
	// ciphertext = ciphertext[8:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	padding := ciphertext[len(ciphertext)-1]
	return ciphertext[:len(ciphertext)-int(padding)], nil
}
