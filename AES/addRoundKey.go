package aes

// 轮密钥加变换
func addRoundKey(a, key *[16]int) {
	for i := 0; i < 16; i++ {
		a[i] ^= key[i]
	}
}
