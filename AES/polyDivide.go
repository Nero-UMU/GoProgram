package aes

// 多项式除法
func polyDivide(a, b int, remainder *int) int {
	var tmp int
	c := numb_bits(a) - numb_bits(b)
	var value int = 0
	for c >= 0 {
		value = (1 << c) | value //计算商，这时计算的c值就是每一个是1的位置，直接与1相与，就可以将其设置为1
		tmp = b
		tmp = tmp << c
		a = a ^ tmp //将被除数向左移c位再与a异或，并赋给a
		c = numb_bits(a) - numb_bits(b)
	}
	*remainder = a
	return value
}

// 计算一个十进制数的二进制表示时的位数有多大
func numb_bits(v int) int {
	count := 0
	for v > 0 {
		v = v >> 1
		count++
	}
	return count
}
