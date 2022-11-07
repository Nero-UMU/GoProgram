package aes

// 求逆元
func exgcd(b int) int {
	return extEuclid(b, 283)
}
