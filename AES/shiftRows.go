package aes

// 行移位
func shiftRows(a *[16]int) {
	b := make([]int, 16)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			b[i*4+j] = a[(i*4+j*5)%16]
		}
	}
	for i := 0; i < 16; i++ {
		a[i] = b[i]
	}
}

// 行移位逆变换
func shiftRowsP_1(a *[16]int) {
	b := make([]int, 16)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			b[i*4+j] = a[(16+i*4-j*3)%16]
		}
	}
	for i := 0; i < 16; i++ {
		a[i] = b[i]
	}
}
