package main

import "fmt"

func main() {
	// LFSR 的初始状态和反馈多项式
	var state uint8 = 0x7F    //1111111
	var feedback uint8 = 0x1D // 11101

	// 生成 8 位随机数序列
	for i := 0; i < 8; i++ {
		// 打印结果
		fmt.Printf("before %08b\n", state)
		// 执行一次移位操作
		lsb := state & 0x01
		state >>= 1

		// 根据反馈多项式计算新的 LSB
		if lsb == 1 {
			state ^= feedback
		}

		// 打印结果
		fmt.Printf("after %08b\n", state)
	}
}
