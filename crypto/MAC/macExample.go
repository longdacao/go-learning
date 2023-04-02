package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mac key message")
		os.Exit(1)
	}

	// 从命令行参数中获取密钥和消息
	key := []byte(os.Args[1])
	message := []byte(os.Args[2])

	// 使用SHA-256散列函数和密钥创建HMAC
	mac := hmac.New(sha256.New, key)

	// 将消息写入HMAC
	mac.Write(message)

	// 生成MAC
	expectedMAC := mac.Sum(nil)

	// 将MAC打印到控制台上
	fmt.Printf("Expected MAC: %x\n", expectedMAC)

	// 验证MAC
	ok := hmac.Equal(expectedMAC, computeMAC(key, message))

	if ok {
		fmt.Println("MAC verification succeeded")
	} else {
		fmt.Println("MAC verification failed")
	}
}

// 独立的 computeMAC 函数，从密钥和消息计算MAC
func computeMAC(key, message []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
