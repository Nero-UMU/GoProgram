package main

import "fmt"

func main() {
	// 不带参数
	a := func() {
		fmt.Println("hello world, 1")
	}
	a()

	func() {
		fmt.Println("hello world, 2")
	}()

	// 带参数
	b := func(str string) {
		fmt.Println(str)
	}
	b("hello world, 3")

	func(str string) {
		fmt.Println(str)
	}("hello world, 4")

	// 带返回值
	c := func(str string) string {
		fmt.Println(str)
		return "fuck me"
	}
	x := c("hello world, 5")
	fmt.Println(x)
}
