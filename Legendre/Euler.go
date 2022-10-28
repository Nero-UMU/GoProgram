package main

// 欧拉判定法
func Euler(a, p int) int {
	// 调用模平方法,计算以a为底,(p-1)/2次幂模p
	return Square_and_Multiply(a, (p-1)/2, p)
}
