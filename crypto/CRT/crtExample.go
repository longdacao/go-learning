package main

import (
	"fmt"
)

// 求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 扩展欧几里得算法，求ax+by=gcd(a,b)的一组整数解(x,y)
func extgcd(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	x, y := extgcd(b, a%b)
	x, y = y, x-(a/b)*y
	return x, y
}

// 求乘法逆元，即ax≡1(mod m)的x值
func inv(a, m int) int {
	x, _ := extgcd(a, m)
	return (x%m + m) % m
}

// CRT算法，求x≡b[i](mod m[i])的最小正整数解x
func crt(b []int, m []int) int {
	n := len(b)
	M := 1 // M是所有模数的乘积
	for i := 0; i < n; i++ {
		M *= m[i]
	}
	res := 0 // res是最终结果
	for i := 0; i < n; i++ {
		t := M / m[i]                  // t是除了m[i]以外其他模数的乘积
		res += b[i] * t * inv(t, m[i]) // 根据中国剩余定理公式计算结果
		res %= M                       // 对M取模，保证结果在[0,M-1]范围内
	}
	return res // 返回最小正整数解
}

func main() {
	b := []int{2, 3}       // 同余方程组右边的常数项数组
	m := []int{3, 5}       // 同余方程组左边的模数数组
	fmt.Println(crt(b, m)) // 输出结果为8，即x≡8(mod 15)
}
