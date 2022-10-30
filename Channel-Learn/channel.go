package main

import "fmt"

func addNum(num int, ch chan int) {
	ch <- num
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	channel := make(chan int)
	go addNum(arr[0], channel)
	go addNum(arr[1], channel)
	go addNum(arr[2], channel)
	go addNum(arr[3], channel)
	go addNum(arr[4], channel)
	go addNum(arr[5], channel)
	go addNum(arr[6], channel)
	go addNum(arr[7], channel)
	a := <-channel
	b := <-channel
	c := <-channel
	d := <-channel
	e := <-channel
	f := <-channel
	g := <-channel
	h := <-channel
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
	fmt.Printf("c=%v\n", c)
	fmt.Printf("d=%v\n", d)
	fmt.Printf("e=%v\n", e)
	fmt.Printf("f=%v\n", f)
	fmt.Printf("g=%v\n", g)
	fmt.Printf("h=%v\n", h)
}
