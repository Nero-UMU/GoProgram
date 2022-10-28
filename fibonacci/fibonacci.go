package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		temp := i
		i, j = j, (i + j)
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
