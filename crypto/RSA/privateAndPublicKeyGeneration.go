package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateRSAKeyPair(bitSize int) (n, e, d *big.Int, err error) {
	// 计算 p 和 q
	p, err := rand.Prime(rand.Reader, bitSize/2)
	fmt.Println("p:", p)
	if err != nil {
		return nil, nil, nil, err
	}
	q, err := rand.Prime(rand.Reader, bitSize/2)
	fmt.Println("q:", q)
	if err != nil {
		return nil, nil, nil, err
	}

	// 计算 n = p*q
	n = new(big.Int).Mul(p, q)

	// 计算 phi(n) = (p-1)*(q-1)
	phiN := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	fmt.Println("phiN:", phiN)

	// 选择一个小于 phi(n) 且和 phi(n) 互质的整数 e
	e = big.NewInt(65537)
	fmt.Println("e before:", e)
	if phiN.Cmp(e) > 0 && e.GCD(nil, nil, phiN, e).Cmp(big.NewInt(1)) != 0 {
		e = e.Add(e, big.NewInt(2))
	}

	fmt.Println("e after:", e)
	// 计算 d = e^(-1) mod phi(n)
	d = new(big.Int)
	if d.ModInverse(e, phiN) == nil {
		return nil, nil, nil, fmt.Errorf("Error computing private exponent")
	}

	return n, e, d, nil
}

func main() {
	n, e, d, err := generateRSAKeyPair(2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("n:", n)
	fmt.Println("e:", e)
	fmt.Println("d:", d)
}
