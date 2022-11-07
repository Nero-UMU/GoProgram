package aes

// AES 加密
func AES(plainText, key *[16]int) *[16]int {
	var roundKey [176]int
	var cipherText [16]int
	keyExpansion(key, &roundKey)
	for i := 0; i < 16; i++ {
		cipherText[i] = plainText[i]
	}
	addRoundKey(&cipherText, (*[16]int)(roundKey[0:16]))
	for round := 1; round < 10; round++ {
		for i := 0; i < 16; i++ {
			cipherText[i] = calcBox(cipherText[i])
		}
		shiftRows(&cipherText)
		mixColumns(&cipherText)
		addRoundKey(&cipherText, (*[16]int)(roundKey[round*16:(round+1)*16]))
	}
	for i := 0; i < 16; i++ {
		cipherText[i] = calcBox(cipherText[i])
	}
	shiftRows(&cipherText)
	addRoundKey(&cipherText, (*[16]int)(roundKey[160:176]))
	return &cipherText
}
