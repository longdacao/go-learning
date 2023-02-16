package main

import "fmt"

func main() {
	// 初始化一个空的切片
	b := []byte{}

	// 添加一个元素
	b = append(b, 'H')

	// 添加多个元素
	b = append(b, 'e', 'l', 'l', 'o')

	fmt.Println(string(b)) // 输出: "Hello"
}
