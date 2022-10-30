package main

import "fmt"

type coder interface {
	one()
	two()
}

type Vertex struct {
	S string
}

func (v Vertex) one() {
	fmt.Printf("I am one\n")
}

func (v *Vertex) two() {
	fmt.Printf("I am two\n")
}

func main() {
	var c coder = &Vertex{"Go"}
	c.one()
	c.two()
}
