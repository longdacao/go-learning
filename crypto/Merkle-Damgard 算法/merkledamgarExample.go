package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// 要散列的原始消息
	message := "hello world"

	// 补位和填充消息
	padded := pad([]byte(message))

	// 分组长度为 64 字节
	blockSize := 64

	// 初始化哈希值
	hash := [32]byte{
		0x6a, 0x09, 0xe6, 0x67,
		0xf3, 0xbc, 0xc9, 0x08,
		0xbb, 0x67, 0xae, 0x85,
		0x84, 0xca, 0xa7, 0x3b,
		0x3b, 0x23, 0x19, 0x86,
		0x8c, 0x5f, 0x5d, 0x51,
		0x6d, 0x36, 0x98, 0x2d,
		0x82, 0x39, 0x2e, 0xf6,
	}

	// 分组循环处理
	for i := 0; i < len(padded); i += blockSize {
		chunk := padded[i : i+blockSize]

		// 哈希散列处理
		hash = sha256.Sum256(append(chunk, hash[:]...))
	}

	// 最终散列值
	finalHash := hex.EncodeToString(hash[:])
	fmt.Println("result =", finalHash)
}

// 补位和填充消息
func pad(message []byte) []byte {
	// 补位
	padded := append(message, byte(0x80))

	// 填充
	for (len(padded)*8)%512 != 448 {
		padded = append(padded, byte(0))
	}

	// 消息长度
	length := uint64(len(message) * 8)
	lengthBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		lengthBytes[i] = byte(length >> uint(56-(i*8)))
	}
	padded = append(padded, lengthBytes...)

	return padded
}
