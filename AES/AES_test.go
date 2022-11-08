package aes

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	var plainText = [16]int{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	var key = [16]int{0x0f, 0x15, 0x71, 0xc9, 0x47, 0xd9, 0xe8, 0x59, 0x0c, 0xb7, 0xad, 0xd6, 0xaf, 0x7f, 0x67, 0x98}
	fmt.Printf("明文    :")
	for i := 0; i < 16; i++ {
		if i%16 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%X ", plainText[i])
	}
	fmt.Printf("\n加密密钥:")
	for i := 0; i < 16; i++ {
		if i%16 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%X ", key[i])
	}
	fmt.Println()
	cipherText := AES_encrypt((*[16]int)(&plainText), (*[16]int)(&key))
	fmt.Printf("加密结果:")
	for i := 0; i < 16; i++ {
		if i%16 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%X ", cipherText[i])
	}

	fmt.Println()
	newPlainText := AES_decrypt(cipherText, (*[16]int)(&key))
	fmt.Printf("解密结果:")
	for i := 0; i < 16; i++ {
		if i%16 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%X ", newPlainText[i])
	}
	fmt.Println()
}
