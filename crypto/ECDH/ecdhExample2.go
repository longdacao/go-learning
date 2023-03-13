package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	// 生成公私钥对
	curve := elliptic.P256()
	privKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err.Error())
	}
	pubKey := append(x.Bytes(), y.Bytes()...)

	// 打印公私钥
	fmt.Println("Private Key:", hex.EncodeToString(privKey))
	fmt.Println("Public Key:", hex.EncodeToString(pubKey))

	// 模拟 Alice 和 Bob 公钥交换
	alicePubKeyX := new(big.Int)
	alicePubKeyY := new(big.Int)
	bobPubKeyX := new(big.Int)
	bobPubKeyY := new(big.Int)

	// Alice 发送公钥给 Bob
	alicePubKey := pubKey
	alicePubKeyX.SetBytes(alicePubKey[:32])
	alicePubKeyY.SetBytes(alicePubKey[32:])

	// Bob 发送公钥给 Alice
	bobPubKey := pubKey
	bobPubKeyX.SetBytes(bobPubKey[:32])
	bobPubKeyY.SetBytes(bobPubKey[32:])

	// 计算共享密钥
	aliceSharedKeyX, _ := curve.ScalarMult(bobPubKeyX, bobPubKeyY, privKey)
	bobSharedKeyX, _ := curve.ScalarMult(alicePubKeyX, alicePubKeyY, privKey)

	if aliceSharedKeyX.Cmp(bobSharedKeyX) == 0 {
		fmt.Println("Shared Key:", hex.EncodeToString(aliceSharedKeyX.Bytes()))
	} else {
		panic("Shared key calculation error")
	}
}
