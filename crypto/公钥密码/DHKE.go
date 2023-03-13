package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 选择一个大素数p和一个原根g
	p, _ := new(big.Int).SetString("88343549029248851119766409633467555065993168617986079533132344008343211837883", 10)
	g := big.NewInt(2)

	// 生成私钥和公钥
	alicePrivateKey, _ := rand.Int(rand.Reader, p)
	alicePublicKey := new(big.Int).Exp(g, alicePrivateKey, p)

	bobPrivateKey, _ := rand.Int(rand.Reader, p)
	bobPublicKey := new(big.Int).Exp(g, bobPrivateKey, p)

	// 交换公钥
	sharedKey1 := new(big.Int).Exp(bobPublicKey, alicePrivateKey, p)
	sharedKey2 := new(big.Int).Exp(alicePublicKey, bobPrivateKey, p)

	fmt.Println("Alice private key:", alicePrivateKey)
	fmt.Println("Alice public key:", alicePublicKey)
	fmt.Println("Bob private key:", bobPrivateKey)
	fmt.Println("Bob public key:", bobPublicKey)
	fmt.Println("Shared key1:", sharedKey1)
	fmt.Println("Shared key2:", sharedKey2)
}
