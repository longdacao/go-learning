package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func FermatTest(n int) string {
	if n == 2 {
		return "Prime"
	}
	if n < 2 || n%2 == 0 {
		return "Composite"
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := big.NewInt(rand.Int63n(int64(n-2)) + 2) // 随机选取 2 到 n-1 之间的整数
		t := big.NewInt(int64(n - 1))
		x := new(big.Int).Exp(a, t, big.NewInt(int64(n)))

		if x.Cmp(big.NewInt(1)) != 0 {
			return "Composite"
		}
	}

	return "Prime"
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if FermatTest(n) == "Prime" {
		fmt.Println(n, "is a prime number.")
	} else {
		fmt.Println(n, "is not a prime number.")
	}
}
