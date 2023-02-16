package main

import "fmt"

// 加密函数，将明文文本移位加密成密文
func caesarEncrypt(plainText string, shift int) string {
	var cipherText string
	for _, r := range plainText {
		if r >= 'A' && r <= 'Z' {
			cipherText += string('A' + (r-'A'+rune(shift))%26)
		} else if r >= 'a' && r <= 'z' {
			cipherText += string('a' + (r-'a'+rune(shift))%26)
		} else {
			cipherText += string(r)
		}
	}
	return cipherText
}

// 解密函数，将密文移位解密成明文文本
func caesarDecrypt(cipherText string, shift int) string {
	var plainText string
	for _, r := range cipherText {
		if r >= 'A' && r <= 'Z' {
			plainText += string('A' + (r-'A'-rune(shift)+26)%26)
		} else if r >= 'a' && r <= 'z' {
			plainText += string('a' + (r-'a'-rune(shift)+26)%26)
		} else {
			plainText += string(r)
		}
	}
	return plainText
}

func main() {
	fmt.Println(caesarEncrypt("ASFeasdfw", 4))
}
