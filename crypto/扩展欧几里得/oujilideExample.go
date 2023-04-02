package main

import "fmt"

func extendedEuclid(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x1, y1 := extendedEuclid(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return gcd, x, y
}

func main() {
	a, b := 15, 35
	gcd, x, y := extendedEuclid(a, b)
	fmt.Printf("gcd(%d, %d) = %d, x = %d, y = %d\n", a, b, gcd, x, y)
}
