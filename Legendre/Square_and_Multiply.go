package main

// 模重复平方法,计算以m为底的e次幂取模n (m^e modn)
func Square_and_Multiply(m, e, n int) int {
	// a存放计算结果，初始化为1
	a := 1
	var k, num int = 0, e
	// k为e的二进制最高位的位数
	for num != 0 {
		num = num >> 1
		k++
	}
	for i := k - 1; i > -1; i-- {
		// 每一轮都平方一次,取模n
		a = a * a % n
		if (e>>i)&1 == 1 {
			// 如果这一轮的二进制数值为1则再乘m取模n
			a = a * m % n
		}
	}
	return a
}
