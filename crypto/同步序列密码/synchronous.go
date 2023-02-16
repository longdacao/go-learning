package main

import (
	"errors"
	"fmt"
)

// 加密函数，使用同步序列密码加密明文文本
func synchronousEncrypt(plainText string, key string) (string, error) {
	if len(key) != len(plainText) {
		return "", errors.New("Length not match")
	}
	var cipherText string
	for i := 0; i < len(plainText); i++ {
		c := plainText[i] ^ key[i]
		cipherText += string(c)
	}
	return cipherText, nil
}

// 解密函数，使用同步序列密码解密密文
func synchronousDecrypt(cipherText string, key string) (string, error) {
	if len(key) != len(cipherText) {
		return "", errors.New("Length not match")
	}

	var plainText string
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] ^ key[i]
		plainText += string(c)
	}
	return plainText, nil
}

func main() {
	var result, errorMessage = synchronousEncrypt("asd", "e3e")
	fmt.Println(result, errorMessage)
	result, errorMessage = synchronousDecrypt(result, "e3e")
	fmt.Println(result, errorMessage)
}
