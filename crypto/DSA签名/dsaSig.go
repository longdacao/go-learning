package main

import (
	"crypto"
	"crypto/dsa"
	"crypto/rand"
	"encoding/asn1"
	"fmt"
)

func main() {
	// 生成DSA密钥对
	privateKey, err := dsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 待签名的数据
	data := []byte("hello, world!")

	// 用私钥对数据进行签名
	sig, err := dsa.SignASN1(rand.Reader, privateKey, crypto.SHA256.Sum(data)[:])
	if err != nil {
		panic(err)
	}

	// 输出签名的字符串格式
	sigStr := fmt.Sprintf("%x", sig)
	fmt.Printf("DSA signature: %s\n", sigStr)

	// 将签名的字符串格式转换为字节数组
	sigBytes, err := decodeHexString(sigStr)
	if err != nil {
		panic(err)
	}

	// 验证签名的有效性
	verified := dsa.VerifyASN1(publicKey, crypto.SHA256.Sum(data)[:], sigBytes)
	if verified {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature is invalid!")
	}
}

// Hex字符串转字节数组
func decodeHexString(str string) ([]byte, error) {
	bytes, err := asn1.Unmarshal([]byte(str), nil)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
