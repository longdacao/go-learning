package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	// 生成一个新的密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// 从私钥中导出公钥
	publicKey := &privateKey.PublicKey

	// 使用 PEM 编码将私钥保存到文件中
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privateKeyPEMBytes := pem.EncodeToMemory(privateKeyPEM)
	fmt.Println(string(privateKeyPEMBytes))

	// 使用 PEM 编码将公钥保存到文件中
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}
	publicKeyPEMBytes := pem.EncodeToMemory(publicKeyPEM)
	fmt.Println(string(publicKeyPEMBytes))

	// 加密和解密消息
	message := []byte("Hello, RSA!")
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plaintext: %s\n", plaintext)
}
