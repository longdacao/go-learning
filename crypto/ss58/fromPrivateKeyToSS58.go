package main

import (
	"fmt"

	"github.com/centrifuge/go-substrate-crypto/ss58"
	"golang.org/x/crypto/ed25519"
)

func main() {
	// 生成密钥对
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	// 打印公钥和私钥
	fmt.Printf("public key: %x\n", publicKey)
	fmt.Printf("private key: %x\n", privateKey)

	// 使用 ss58 编码将公钥转换为 substrate 网络地址
	substrateAddress := ss58.Encode(publicKey, 42)
	fmt.Printf("substrate address: %s\n", substrateAddress)

	// 使用私钥签名
	message := []byte("Hello, World!")
	signature := ed25519.Sign(privateKey, message)

	// 打印签名结果
	fmt.Printf("signature: %x\n", signature)

	// 使用 ss58 编码将签名转换为 substrate 网络地址
	substrateSignature := ss58.Encode(signature, 42)
	fmt.Printf("substrate signature: %s\n", substrateSignature)
}
