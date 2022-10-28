package main

import "fmt"

type MyInt interface {
	fun()
}

type vertex struct {
	S string
}

type MyFloat float64

func (v vertex) fun() {
	fmt.Printf("用vertex实现了方法\n")
	fmt.Println(v)
}

func (f MyFloat) fun() {
	fmt.Printf("用MyFloat实现了方法\n")
	fmt.Println(f)
}

func main() {
	var i MyInt
	i = vertex{"hehe"}
	i.fun()
	var j MyInt = MyFloat(20)
	j.fun()
}
