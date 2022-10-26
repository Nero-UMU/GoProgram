package main

import "fmt"

func main() {
	var i int
	var p *int
	i = 20
	p = &i
	fmt.Printf("type is %T, value is %v and point to %v \n", p, p, *p)
	*p = 100
	fmt.Printf("changed i value is %v\n", i)
}
