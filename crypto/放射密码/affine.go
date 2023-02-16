package main

import "fmt"

// 加密函数，将明文文本使用仿射密码加密成密文
func affineEncrypt(plainText string, a, b int) string {
	var cipherText string
	for _, r := range plainText {
		if r >= 'A' && r <= 'Z' {
			cipherText += string('A' + ((int(r)-'A')*a+b)%26)
		} else if r >= 'a' && r <= 'z' {
			cipherText += string('a' + ((int(r)-'a')*a+b)%26)
		} else {
			cipherText += string(r)
		}
	}
	return cipherText
}

// 解密函数，将密文使用仿射密码解密成明文文本
func affineDecrypt(cipherText string, a, b int) string {
	var plainText string
	for _, r := range cipherText {
		if r >= 'A' && r <= 'Z' {
			plainText += string('A' + (((int(r)-'A'-b+26)%26)*modInverse(a, 26))%26)
		} else if r >= 'a' && r <= 'z' {
			plainText += string('a' + (((int(r)-'a'-b+26)%26)*modInverse(a, 26))%26)
		} else {
			plainText += string(r)
		}
	}
	return plainText
}

// 求逆元函数，用于求解a的逆元
func modInverse(a, m int) int {
	a %= m
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return 1
}

func main() {
	fmt.Println(affineEncrypt("azasdff", 3, 6))
	fmt.Println(affineDecrypt("gdgipvv", 3, 6))
}
