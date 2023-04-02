package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	// 生成 RSA 密钥对
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 加密数据
	plaintext := []byte("Hello, world!")
	ciphertext, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privKey.PublicKey, plaintext, nil)

	// 解密数据
	// 先计算出 p 和 q
	p := privKey.Primes[0]
	q := privKey.Primes[1]

	// 计算 dp 和 dq
	dp := privKey.D.Mod(privKey.D, p.Sub(p, big.NewInt(1)))
	dq := privKey.D.Mod(privKey.D, q.Sub(q, big.NewInt(1)))

	// 计算 qInv，即 q 在模 p 下的逆元
	qInv := new(big.Int).ModInverse(q, p)

	// 使用 CRT 解密
	cipherBigInt := new(big.Int).SetBytes(ciphertext)
	mp := new(big.Int).Exp(cipherBigInt, dp, p)
	mq := new(big.Int).Exp(cipherBigInt, dq, q)
	h := new(big.Int).Mul(qInv, new(big.Int).Sub(mp, mq))
	m := new(big.Int).Add(mq, new(big.Int).Mul(h, q))

	// 打印解密结果
	decrypted := m.Bytes()
	fmt.Println(string(decrypted))
}
