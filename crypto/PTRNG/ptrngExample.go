package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100)

	fmt.Println(randomNumber)
}
