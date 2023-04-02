package main

import "fmt"

func phi(n int) int {
	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func main() {
	n := 12
	fmt.Printf("phi(%d) = %d\n", n, phi(n))
}
