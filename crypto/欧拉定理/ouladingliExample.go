package main

import (
	"fmt"
	"math/big"
)

// 计算 a 的 b 次方对 m 取模的值
func modExp(a, b, m *big.Int) *big.Int {
	if m.Sign() == 0 {
		return nil // 模数不能为 负数
	}
	res := big.NewInt(1)
	a = a.Mod(a, m)
	for b.Sign() > 0 {
		if b.Bit(0) == 1 {
			res = res.Mul(res, a).Mod(res, m)
		}
		b.Rsh(b, 1)
		a = a.Mul(a, a).Mod(a, m)
	}
	return res
}

// 计算欧拉函数
func eulerPhi(n *big.Int) *big.Int {
	if n.Sign() == 0 {
		return big.NewInt(0) // 对于 0，欧拉函数的值为 0
	}
	res := new(big.Int).Set(n) // 将 res 初始化为 n
	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		if i.ProbablyPrime(10) && new(big.Int).Mod(n, i).Sign() == 0 {
			// 若 i 是 n 的质因子，则将 res 除以 i，同时在循环中除去所有 i 的倍数
			res.Div(res, i)
			res.Mul(res, new(big.Int).Sub(i, big.NewInt(1)))
			for new(big.Int).Mod(n, i).Sign() == 0 {
				n.Div(n, i)
			}
		}
	}
	return res
}

func main() {
	n := big.NewInt(15)
	fmt.Println("欧拉函数值：", eulerPhi(n))
	fmt.Println("欧拉定理值：", modExp(big.NewInt(7), eulerPhi(n), n))
}
