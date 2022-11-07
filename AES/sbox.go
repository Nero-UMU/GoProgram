package aes

// 计算S盒
func calcBox(want int) int {
	want = exgcd(want)
	want = tranceFormP(want)
	return want
}

// 计算逆S盒
func calcExtBox(want int) int {
	want = tranceFormP_1(want)
	want = exgcd(want)
	return want
}

// S盒位变换
func tranceFormP(num int) int {
	var result int = 0
	var n = make([]int, 8)
	var b = make([]int, 8)
	var c = [8]int{1, 1, 0, 0, 0, 1, 1, 0}
	for i := 0; i < 8; i++ {
		n[i] = num % 2
		num = num >> 1
	}
	j := 1
	for i := 0; i < 8; i++ {
		b[i] = n[i] ^ n[(i+4)%8] ^ n[(i+5)%8] ^ n[(i+6)%8] ^ n[(i+7)%8] ^ c[i]
		result += b[i] * j
		j *= 2
	}
	return result
}

// 逆S盒位变换
func tranceFormP_1(num int) int {
	var result int = 0
	var n = make([]int, 8)
	var b = make([]int, 8)
	var d = [8]int{1, 0, 1, 0, 0, 0, 0, 0}
	for i := 0; i < 8; i++ {
		n[i] = num % 2
		num = num >> 1
	}
	j := 1
	for i := 0; i < 8; i++ {
		b[i] = n[(i+2)%8] ^ n[(i+5)%8] ^ n[(i+7)%8] ^ d[i]
		result += b[i] * j
		j *= 2
	}
	return result
}
