package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 生成公私钥对
	bits := 1024
	p, _ := paillier.GenerateKeyPair(bits)

	// 加密 10 和 5
	c10, _ := p.Encrypt(big.NewInt(10))
	c5, _ := p.Encrypt(big.NewInt(5))

	// 解密 c10 和 c5 并输出
	m10, _ := p.Decrypt(c10)
	m5, _ := p.Decrypt(c5)
	fmt.Printf("c10 = %v, decrypted = %v\n", c10, m10)
	fmt.Printf("c5 = %v, decrypted = %v\n", c5, m5)

	// 计算同态加法 C = c10 * c5 mod N^2
	cSum := p.Add(c10, c5)
	mSum, _ := p.Decrypt(cSum)
	fmt.Printf("cSum = %v, decrypted = %v\n", cSum, mSum)

	// 计算同态乘法 C = c10^3 mod N^2
	cProduct := p.Exp(c10, 3)
	mProduct, _ := p.Decrypt(cProduct)
	fmt.Printf("cProduct = %v, decrypted = %v\n", cProduct, mProduct)
}
