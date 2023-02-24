package main

import (
	"bytes"
	"fmt"
	"math/big"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base58Encode encodes a byte slice to a Base58 string
func Base58Encode(input []byte) string {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(alphabet)))
	zero := big.NewInt(0)
	for x.Cmp(zero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, base, mod)
		// fmt.Println("x======", x)
		// fmt.Println("mod======", mod)
		// fmt.Println("alphabet[mod.Int64()]======", alphabet[mod.Int64()])
		result = append(result, alphabet[mod.Int64()])
	}

	// fmt.Println("before swap result======", result)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Base58Decode decodes a Base58 string to a byte slice
func Base58Decode(input string) []byte {
	x := big.NewInt(0)
	base := big.NewInt(int64(len(alphabet)))
	for _, r := range input {
		fmt.Println("r======", r)
		index := bytes.IndexByte([]byte(alphabet), byte(r))
		fmt.Println("index======", index)
		if index < 0 {
			return nil
		}
		x.Mul(x, base)
		x.Add(x, big.NewInt(int64(index)))
	}
	fmt.Println("Decode x======", x)
	result := x.Bytes()
	fmt.Println("Decode result======", result)
	result = append(bytes.Repeat([]byte{0x00}, len(input)-len(result)), result...)
	return result
}

func main() {
	result := Base58Encode([]byte{0xff, 0xff})
	fmt.Println("final result======", result)
	result2 := Base58Decode(result)
	fmt.Println("final result2======", result2)
	// fmt.Println("result======", big.NewInt(0).SetBytes([]byte{0xff, 0xff}))
	// fmt.Println("result======", big.NewInt(0))
	// str := "hello"
	// for i, c := range str {
	// 	fmt.Printf("字符 %c 的下标为 %d\n", c, i)
	// }
}
