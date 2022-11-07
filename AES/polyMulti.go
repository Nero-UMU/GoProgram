package aes

// 多项式乘法
func polyMulti(a, b int) int {
	value := a
	result := a
	var r int
	if a == 0 || b == 0 {
		return 0
	}
	r = b % 2
	b = b >> 1
	if r == 0 { //判断最后一位是0还是1,是0的话result就等于0
		result = 0
	}
	for b != 0 {
		r = b % 2
		b = b >> 1
		if (value>>7)&1 == 1 { // b7为1,就要与00011011异或
			value = (value << 1) ^ 0x1b
		} else {
			value = value << 1
		}
		if r == 1 { // 此位为1的结果就参加异或运算
			result = result ^ value
		}
	}
	return result
}
