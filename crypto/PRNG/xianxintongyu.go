package main

import (
	"fmt"
)

// 定义线性同余 PRNG 发生器结构体
type LCG struct {
	seed uint64 // 种子
	a    uint64 // 增量
	c    uint64 // 偏移量
	m    uint64 // 模数
}

// 创建一个新的 LCG 发生器
func NewLCG(seed, a, c, m uint64) *LCG {
	return &LCG{seed: seed, a: a, c: c, m: m}
}

// 生成下一个随机数
func (lcg *LCG) Next() uint64 {
	lcg.seed = (lcg.a*lcg.seed + lcg.c) % lcg.m
	return lcg.seed
}

func main() {
	// 创建一个 LCG 发生器
	lcg := NewLCG(1234, 1103515245, 12345, 1<<31)

	// 生成 10 个随机数
	for i := 0; i < 10; i++ {
		// 生成下一个随机数并将其打印到控制台
		fmt.Println(lcg.Next())
	}
}
