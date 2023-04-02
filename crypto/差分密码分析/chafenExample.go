package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func main() {
	// 设置要破解的密文和明文
	ciphertext := "3fa40e8a984d48156a271787ab8883f9893d51ec4b563b53" +
		"05769334be5a6248141a5d1ba3a8a0f5c1d22d10f44e6e4f" +
		"11a49f55c2d835bfb0f7033c1a3cdc2c2dd91896d6b6a1e7"
	plaintext := "0123456789abcdef"

	// 将密文和明文解码为字节数组
	ct, _ := hex.DecodeString(ciphertext)
	fmt.Println("ct======", ct)
	pt, _ := hex.DecodeString(plaintext)
	fmt.Println("pt=======", pt)

	// 循环猜测密钥的每个位
	for i := 0; i < 8; i++ {
		// 生成差分明文和差分密文
		diffPt := make([]byte, 8)
		diffCt := make([]byte, 8)
		diffPt[i] = 0x80
		fmt.Println("diffPt=======", diffPt)
		diffCt[i] = ct[i] ^ pt[i] ^ 0x80

		// 尝试每个可能的密钥
		for j := 0; j < 256; j++ {
			key := make([]byte, 8)
			key[i] = byte(j)

			// 创建DES密码块
			block, err := des.NewCipher(key)
			if err != nil {
				panic(err)
			}

			// 创建CBC模式密码分组
			mode := cipher.NewCBCDecrypter(block, make([]byte, 8))

			// 对差分明文进行解密，得到差分密文
			mode.CryptBlocks(diffCt, diffPt)

			// 比较得到的差分密文和实际的差分密文
			if diffCt[i] == ct[i]^pt[i]^0x80 {
				fmt.Printf("Key found: %x\n", key)
			}
		}
	}
}
