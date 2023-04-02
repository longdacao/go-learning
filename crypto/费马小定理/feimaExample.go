package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// 模幂运算 a^b mod m
func modExp(a, b, m *big.Int) *big.Int {
	res := big.NewInt(1)
	for b.BitLen() > 0 {
		if b.Bit(0) == 1 {
			res.Mul(res, a).Mod(res, m)
		}
		a.Mul(a, a).Mod(a, m)
		b.Rsh(b, 1)
	}
	return res
}

// 费马素性测试
func fermatTest(n *big.Int, k int) bool {
	if n.BitLen() == 1 {
		return false // 排除 0 和 1 的情况
	}
	if n.Bit(0) == 0 {
		return n.Cmp(big.NewInt(2)) == 0 // 排除偶数
	}

	// 进行 k 次测试
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(2)))
		a.Add(a, big.NewInt(2)) // 生成 [2, n-1] 范围内的随机数
		if modExp(a, new(big.Int).Sub(n, big.NewInt(1)), n).Cmp(big.NewInt(1)) != 0 {
			return false // 如果不通过测试则判断 n 为合数
		}
	}
	return true // 经过 k 次测试后，n 很有可能为素数
}

func main() {
	n := big.NewInt(101)
	k := 10
	if fermatTest(n, k) {
		fmt.Printf("%v is probably a prime number\n", n)
	} else {
		fmt.Printf("%v is composite\n", n)
	}
}
