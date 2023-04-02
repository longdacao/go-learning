package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 生成随机数模n的二次剩余
func getQuadResidue(n *big.Int) (residue *big.Int, err error) {
	for {
		r, err := rand.Int(rand.Reader, n)
		if err != nil {
			return nil, err
		}
		g := new(big.Int).Exp(r, big.NewInt(2), n)
		if g.Cmp(big.NewInt(1)) == 0 {
			continue
		}
		return g, nil
	}
}

// 加密明文m
func encrypt(m, n, g, r *big.Int) (c *big.Int) {
	return new(big.Int).Mul(new(big.Int).Exp(g, m, n), new(big.Int).Exp(r, n, n))
}

// 解密密文c,私钥为L(x-d)/n
func decrypt(c, n, p, q, d *big.Int) *big.Int {
	L := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	mu := new(big.Int).ModInverse(L.Mod(new(big.Int).Mul(L, d), n), n)
	t1 := new(big.Int).Exp(c.Mod(n), new(big.Int).Sub(n, d), n)
	t2 := new(big.Int).Mod(new(big.Int).Mul(t1.Mod(n), mu), n)
	return new(big.Int).Exp(t2.Mod(n), L.Div(L, new(big.Int).Sub(p, big.NewInt(1))), n)
}

func main() {
	// 选取两个大素数p和q
	p, _ := rand.Prime(rand.Reader, 128)
	q, _ := rand.Prime(rand.Reader, 128)

	// 计算n=pq和lcm=(p-1)*(q-1)
	n := new(big.Int).Mul(p, q)
	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))
	lcm := new(big.Int).Mul(p, q)

	// 选择一个随机整数r，使得1 < r < n且gcd(r,n) = 1
	for {
		r, err := rand.Int(rand.Reader, n)
		if err != nil {
			fmt.Println("随机数生成错误：", err)
		}
		if r.Cmp(big.NewInt(1)) == 1 && r.Cmp(n) == -1 && new(big.Int).GCD(nil, nil, r, n).Cmp(big.NewInt(1)) == 0 {
			// 生成随机数模n的二次剩余
			g, err := getQuadResidue(n)
			if err != nil {
				fmt.Println("随机数生成错误：", err)
			}

			// 加密明文m
			m := big.NewInt(10)
			c := encrypt(m, n, g, r)

			// 解密密文c
			d := big.NewInt(4) // 私钥为4
			fmt.Println("明文：", decrypt(c, n, p, q, d))
			break
		}
	}
}
