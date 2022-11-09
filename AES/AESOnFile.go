package aes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile(filePath *string) (int, []int) {
	file, err := os.OpenFile(*filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	var plainText []int
	buf := make([]byte, 16)
	var i int = 0
	for {
		length, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("文件读取失败")
				return 0, plainText
			}
		}
		for j := 0; j < length; j++ {
			plainText = append(plainText, int(buf[j]))
		}
		i++
	}

	return i, plainText
}

func getKey(filePath *string) [16]int {
	file, err := os.OpenFile(*filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	var key [16]int
	buf := make([]byte, 16)
	length, err := file.Read(buf)
	for j := 0; j < length; j++ {
		key[j] = int(buf[j])
	}
	return key
}

func writeAppend(filePath string, text *[16]int) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < 16; i++ {
		if byte(text[i]) == 0 {
			continue
		}
		write.WriteByte(byte(text[i]))
	}
	write.Flush()
}

func encryptFile(filePath *string, key *[16]int) {
	num, plainText := readFile(filePath)
	newPath := strings.Split(*filePath, "/")
	path := "/"
	for i := 0; i < len(newPath)-1; i++ {
		path = path + newPath[i] + "/"
	}
	for i := 0; i < num; i++ {
		cipherText := AES_encrypt((*[16]int)(plainText[i*16:(i+1)*16]), (*[16]int)(key))
		writeAppend(path+"/cipher.txt", cipherText)
	}
}

func decryptFile(filePath *string, key *[16]int) {
	num, cipherText := readFile(filePath)
	newPath := strings.Split(*filePath, "/")
	path := "/"
	for i := 0; i < len(newPath)-1; i++ {
		path = path + newPath[i] + "/"
	}
	for i := 0; i < num; i++ {
		cipherText := AES_decrypt((*[16]int)(cipherText[i*16:(i+1)*16]), (*[16]int)(key))
		writeAppend(path+"/plain.txt", cipherText)
	}
}
