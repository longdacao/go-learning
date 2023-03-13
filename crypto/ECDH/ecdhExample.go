package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	// 选择椭圆曲线
	curve := elliptic.P256()

	// 生成私钥
	privateKeyA, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("privateKeyA")
	fmt.Println(privateKeyA)

	// 生成公钥
	publicKeyA := privateKeyA.PublicKey
	fmt.Println("publicKeyA")
	fmt.Println(publicKeyA)

	// 生成私钥
	privateKeyB, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// 生成公钥
	publicKeyB := privateKeyB.PublicKey

	// Alice 发送公钥给 Bob
	sharedKeyA, err := generateSharedKey(privateKeyA, &publicKeyB)
	if err != nil {
		panic(err)
	}

	// Bob 发送公钥给 Alice
	sharedKeyB, err := generateSharedKey(privateKeyB, &publicKeyA)
	if err != nil {
		panic(err)
	}

	// Alice 和 Bob 计算出来的共享密钥应该是一样的
	fmt.Println("Shared key A:", sharedKeyA)
	fmt.Println("Shared key B:", sharedKeyB)
}

func generateSharedKey(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	// 生成共享密钥
	x, _ := privateKey.PublicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	if x == nil {
		return nil, fmt.Errorf("failed to generate shared key")
	}

	// 将 x 转成字节数组返回
	return x.Bytes(), nil
}
