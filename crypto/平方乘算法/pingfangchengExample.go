package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// 通过平方乘算法计算 x^y mod m
func squareAndMultiply(x, y, m *big.Int) *big.Int {
	// 将 y 转换为二进制
	binY := fmt.Sprintf("%b", y)

	// 用 r 存储计算过程中的值
	r := big.NewInt(1)

	// 用 v 存储中间结果
	v := new(big.Int).Set(x)

	// 遍历 y 的二进制位，从高位到低位
	for i := len(binY) - 1; i >= 0; i-- {
		// 平方
		r.Mul(r, r)
		r.Mod(r, m)

		// 如果当前二进制位是 1，则乘上 x
		if binY[i] == '1' {
			r.Mul(r, v)
			r.Mod(r, m)
		}
	}

	return r
}

func squareAndMultiply2(base, exponent, modulus *big.Int) *big.Int {
	result := big.NewInt(1)

	// Convert exponent to binary
	binaryExp := strconv.FormatInt(int64(exponent.Int64()), 2)
	fmt.Println("binaryExp====", binaryExp)

	// Square and multiply algorithm
	for i := len(binaryExp) - 1; i >= 0; i-- {
		fmt.Println("i====", i)
		fmt.Println("string(binaryExp[i])====", string(binaryExp[i]))
		if string(binaryExp[i]) == "1" {
			result.Mul(result, base)
			// result.Mod(result, modulus)
			fmt.Println("result====", result)
		}
		base.Mul(base, base)
		// base.Mod(base, modulus)
	}

	return result
}

func main() {
	x := big.NewInt(7)
	y := big.NewInt(2)
	m := big.NewInt(100)

	result := squareAndMultiply2(x, y, m)
	fmt.Printf("%v^%v mod %v = %v\n", x, y, m, result)
}
