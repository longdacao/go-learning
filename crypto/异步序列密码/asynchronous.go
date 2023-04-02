package main

import "fmt"

// 加密函数，使用异步序列密码加密明文文本
func asynchronousEncrypt(plainText string, key uint32, length int) string {
	var cipherText []byte
	var shiftReg uint32 = key
	for i := 0; i < len(plainText); i++ {
		// 生成伪随机序列
		var pseudoRandom uint32 = (shiftReg >> 0) ^ (shiftReg >> 2) ^ (shiftReg >> 3) ^ (shiftReg >> 5)
		// 从伪随机序列中取出一个字节
		var c byte = byte(pseudoRandom & 0xff)
		// 将明文中的一个字节与伪随机序列中的一个字节进行异或运算
		c = c ^ plainText[i]
		cipherText = append(cipherText, c)
		// 移位寄存器进行移位，反馈
		shiftReg = ((shiftReg >> 1) | (uint32(c) << 31)) ^ (uint32(c) << 8)
	}
	return string(cipherText)
}

// 解密函数，使用异步序列密码解密密文
func asynchronousDecrypt(cipherText string, key uint32, length int) string {
	var plainText []byte
	var shiftReg uint32 = key
	for i := 0; i < len(cipherText); i++ {
		// 生成伪随机序列
		var pseudoRandom uint32 = (shiftReg >> 0) ^ (shiftReg >> 2) ^ (shiftReg >> 3) ^ (shiftReg >> 5)
		// 从伪随机序列中取出一个字节
		var c byte = byte(pseudoRandom & 0xff)
		fmt.Println("c", c)
		// 将密文中的一个字节与伪随机序列中的一个字节进行异或运算
		c = c ^ cipherText[i]
		plainText = append(plainText, c)
		// 移位寄存器进行移位，反馈
		shiftReg = ((shiftReg >> 1) | (uint32(cipherText[i]) << 31)) ^ (uint32(cipherText[i]) << 8)
	}
	return string(plainText)
}

func main() {
	var result = asynchronousEncrypt("asdfsd", 4568, 3)
	fmt.Println(asynchronousDecrypt(result, 4568, 3))
}
