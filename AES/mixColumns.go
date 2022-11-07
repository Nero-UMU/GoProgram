package aes

// 列混淆
func mixColumns(a *[16]int) {
	b := make([]int, 16)
	for j := 0; j < 4; j++ {
		b[4*j] = (polyMulti(2, a[j*4]) ^ polyMulti(3, a[1+j*4]) ^ a[2+j*4] ^ a[3+j*4]) % 256
	}
	for j := 0; j < 4; j++ {
		b[4*j+1] = (a[j*4] ^ polyMulti(2, a[1+j*4]) ^ polyMulti(3, a[2+j*4]) ^ a[3+j*4]) % 256
	}
	for j := 0; j < 4; j++ {
		b[4*j+2] = (a[j*4] ^ a[1+j*4] ^ polyMulti(2, a[2+j*4]) ^ polyMulti(3, a[3+j*4])) % 256
	}
	for j := 0; j < 4; j++ {
		b[4*j+3] = (polyMulti(3, a[j*4]) ^ a[1+j*4] ^ a[2+j*4] ^ polyMulti(2, a[3+j*4])) % 256
	}
	for i := 0; i < 16; i++ {
		a[i] = b[i]
	}
}

// 逆列混淆
func mixColumnsP_1(a *[16]int) {
	b := make([]int, 16)
	for j := 0; j < 4; j++ {
		b[j] = (polyMulti(0x0e, a[j*4]) ^ polyMulti(0x0b, a[1+j*4]) ^ polyMulti(0x0d, a[2+j*4]) ^ polyMulti(0x09, a[3+j*4])) % 256
	}
	for j := 0; j < 4; j++ {
		b[4*j+1] = (polyMulti(0x09, a[j*4]) ^ polyMulti(0x0e, a[1+j*4]) ^ polyMulti(0x0b, a[2+j*4]) ^ polyMulti(0x0d, a[3+j*4])) % 256
	}
	for j := 0; j < 4; j++ {
		b[4*j+2] = (polyMulti(0x0d, a[j*4]) ^ polyMulti(0x09, a[1+j*4]) ^ polyMulti(0x0e, a[2+j*4]) ^ polyMulti(0x0b, a[3+j*4])) % 256
	}
	for j := 0; j < 4; j++ {
		b[4*j+3] = (polyMulti(0x0b, a[j*4]) ^ polyMulti(0x0d, a[1+j*4]) ^ polyMulti(0x09, a[2+j*4]) ^ polyMulti(0x0e, a[3+j*4])) % 256
	}
	for i := 0; i < 16; i++ {
		a[i] = b[i]
	}
}
