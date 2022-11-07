package aes

// 扩展欧几里得
func extEuclid(a int, b int) int {
	var tmp int
	if a < b {
		tmp = a
		a = b
		b = tmp
	}
	var x0, y0, x1, y1 int = 1, 0, 0, 1
	var tmp1, tmp2, q1 int
	r1 := a
	r2 := b

	i := 1
	//下面直接用扩展欧几里德算法来做
	for r2 != 0 {
		tmp1 = r2
		q1 = polyDivide(r1, r2, &r2) //r1对r2进行多项式除法，商存在q1,余数存在r2
		r1 = tmp1                    //这步比较重要，就是r1要变成上一次除法的被除数
		tmp1 = x1
		tmp2 = y1
		x1 = x0 ^ polyMulti(q1, x1) //计算v,w;每一个v(x)*f(x)+w(x)*g(x)=1成立
		y1 = y0 ^ polyMulti(q1, y1)
		x0 = tmp1
		y0 = tmp2
		i++
	}
	return y0
}
