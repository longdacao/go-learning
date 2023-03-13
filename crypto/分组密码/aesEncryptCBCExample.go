package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	// key 和 plaintext 必须是 16 字节长度的 byte 数组
	key := []byte("0123456789abcdef")
	plaintext := []byte("Hello, World!")

	// 创建 AES 加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 初始化 CBC 模式
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCEncrypter(block, iv)

	// 填充明文
	padded := pad(plaintext, aes.BlockSize)
	// fmt.Println("aes.BlockSize", aes.BlockSize)

	// 加密
	cipherText := make([]byte, len(padded))
	mode.CryptBlocks(cipherText, padded)

	fmt.Printf("Cipher text: %x\n", cipherText)

	// 解密
	mode = cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	// 去除填充
	plainText = unpad(plainText)

	fmt.Printf("Plain text: %s\n", plainText)
}

// pad 填充函数
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// unpad 去除填充函数
func unpad(data []byte) []byte {
	length := len(data)
	padding := int(data[length-1])
	return data[:(length - padding)]
}
