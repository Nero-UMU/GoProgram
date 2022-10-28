package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (V Vertex) Multiply() int {
	return V.X * V.Y
}

func (V Vertex) Plus() int {
	return V.X + V.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v.Multiply()=", v.Multiply())
	fmt.Println("v.Plus()=", v.Plus())
}
