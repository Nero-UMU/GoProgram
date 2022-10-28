package main

import "fmt"

func main() {
	a := 127
	p := 233
	num := Legendre(a, p)
	if num == 1 {
		fmt.Printf("%v是%v的平方剩余, (a/p)=1\n", a, p)
	} else if num == -1 {
		fmt.Printf("%v是%v的平方非剩余, (a/p)=-1\n", a, p)
	} else {
		fmt.Printf("%v|%v, (a/p)=0\n", p, a)
	}
}
