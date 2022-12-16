package main

import (
	"flag"
	"fmt"
)

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

func exGcd(a, b int) int {
	result := 1
	n1, n2, a2, b2, b1 := a, b, 0, 1, 0
	var q, r, t int
	q = n1 / n2
	r = n1 - q*n2
	for r != 0 {
		q = n1 / n2
		r = n1 - q*n2
		n1 = n2
		n2 = r
		t = a2
		a2 = result - q*a2
		result = t
		t = b2
		b2 = b1 - q*b2
		b1 = t
	}
	result = (result + 5*b) % b
	return result
}

// m明文,k随机整数,d对方密钥,a素数,p素数
func elgamal_en(m, k, d, a, p int) (int, int) {
	c1 := Square_and_Multiply(a, k, p)
	b := Square_and_Multiply(a, d, p)
	c2 := (m % p * Square_and_Multiply(b, k, p)) % p
	return c1, c2
}

func elgamal_de(c1, c2, d, p int) int {
	c11 := exGcd(Square_and_Multiply(c1, d, p), p)
	m := (c2 * c11) % p
	return m
}

var m = flag.Int("m", 0, "消息m")
var k = flag.Int("k", 0, "随机整数k")
var d = flag.Int("d", 0, "对方密钥d")
var a = flag.Int("a", 0, "素数a")
var p = flag.Int("p", 0, "模p")

func main() {
	flag.Parse()
	c1, c2 := elgamal_en(*m, *k, *d, *a, *p)
	fmt.Printf("c1=%v,c2=%v\n", c1, c2)
	plain := elgamal_de(c1, c2, *d, *p)
	fmt.Printf("解密后明文:%v\n", plain)
}
