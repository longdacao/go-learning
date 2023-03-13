package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// 判断 n 是否为质数
func MillerRabin(n *big.Int, k int) bool {
	// 如果 n 是偶数或小于2，则不是质数
	if n.Cmp(big.NewInt(2)) < 0 || n.Bit(0) == 0 {
		return false
	}

	// 分解 n - 1 为 d * 2^r
	d, r := new(big.Int).Sub(n, big.NewInt(1)), 0
	for ; d.Bit(0) == 0; r++ {
		d.Rsh(d, 1)
	}

	// 进行 k 次测试
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		// 选取随机整数 a：2 <= a <= n - 2
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(4)))
		a.Add(a, big.NewInt(2))

		// 计算 a^d mod n
		x := new(big.Int).Exp(a, d, n)

		if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
			continue
		}

		for j := 0; j < r-1; j++ {
			x.Exp(x, big.NewInt(2), n)
			if x.Cmp(big.NewInt(1)) == 0 {
				return false
			}
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
				break
			}
		}

		if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			return false
		}
	}

	return true
}

func main() {
	var n big.Int
	var k int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Print("Enter the number of iterations(k): ")
	fmt.Scan(&k)

	if MillerRabin(&n, k) {
		fmt.Println(n.String(), "is probably prime.")
	} else {
		fmt.Println(n.String(), "is composite.")
	}
}
