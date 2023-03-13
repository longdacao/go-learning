package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// generateKeys 生成 Elgamal 的公钥和私钥
func generateKeys() (publicKey, privateKey *big.Int) {
	p, _ := rand.Prime(rand.Reader, 128)
	g := new(big.Int).Add(p, big.NewInt(-1))
	x, _ := rand.Int(rand.Reader, g)
	y := new(big.Int).Exp(big.NewInt(2), x, p)
	return y, x
}

// encrypt 使用 Elgamal 加密明文
func encrypt(plaintext []byte, publicKey *big.Int) (ciphertext []*big.Int, err error) {
	p, _ := rand.Prime(rand.Reader, 128)
	g := new(big.Int).Add(p, big.NewInt(-1))
	for {
		k, _ := rand.Int(rand.Reader, g)
		// 检查 k 的逆元是否存在
		if new(big.Int).GCD(nil, nil, k, p).Cmp(big.NewInt(1)) == 0 {
			// 计算 c1 和 c2
			c1 := new(big.Int).Exp(big.NewInt(2), k, p)
			c2 := new(big.Int).Exp(publicKey, k, p)
			c2.Mul(c2, new(big.Int).SetBytes(plaintext))
			c2.Mod(c2, p)

			return []*big.Int{c1, c2}, nil
		}
	}
}

// decrypt 使用 Elgamal 解密密文
func decrypt(ciphertext []*big.Int, privateKey *big.Int) (plaintext []byte, err error) {
	p := ciphertext[0]
	c2 := ciphertext[1]

	// 计算 c2 的逆元
	c2Inv := new(big.Int).ModInverse(c2, p)

	// 计算明文
	m := new(big.Int).Exp(c2Inv, privateKey, p)
	m.Mul(m, ciphertext[0])
	m.Mod(m, p)

	return m.Bytes(), nil
}

func main() {
	// 生成公钥和私钥
	publicKey, privateKey := generateKeys()
	fmt.Printf("Public key: %v\n", publicKey)
	fmt.Printf("Private key: %v\n", privateKey)

	// 加密明文
	plaintext := []byte("hello")
	ciphertext, _ := encrypt(plaintext, publicKey)
	fmt.Printf("Ciphertext: %v\n", ciphertext)

	// 解密密文
	decrypted, _ := decrypt(ciphertext, privateKey)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
