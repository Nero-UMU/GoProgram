package aes

// 行移位
func shiftRows(a *[16]int) {
	b := make([]int, 16)
	b[0] = a[0]
	b[4] = a[4]
	b[8] = a[8]
	b[12] = a[12]
	b[1] = a[5]
	b[5] = a[9]
	b[9] = a[13]
	b[13] = a[1]
	b[2] = a[10]
	b[6] = a[14]
	b[10] = a[2]
	b[14] = a[6]
	b[3] = a[15]
	b[7] = a[3]
	b[11] = a[7]
	b[15] = a[11]
	for i := 0; i < 16; i++ {
		a[i] = b[i]
	}
}

// 行移位逆变换
func shiftRowsP_1(a *[16]int) {
	b := make([]int, 16)
	b[0] = a[0]
	b[4] = a[4]
	b[8] = a[8]
	b[12] = a[12]
	b[1] = a[13]
	b[5] = a[1]
	b[9] = a[5]
	b[13] = a[9]
	b[2] = a[10]
	b[6] = a[14]
	b[10] = a[2]
	b[14] = a[6]
	b[3] = a[7]
	b[7] = a[11]
	b[11] = a[15]
	b[15] = a[3]
	for i := 0; i < 16; i++ {
		a[i] = b[i]
	}
}
