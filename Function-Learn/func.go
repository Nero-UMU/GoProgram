package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (V Vertex) Multiply() int {
	return V.X * V.Y
}

func (V *Vertex) Trans(i int) {
	V.X = V.X + i
	V.Y = V.Y - i
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v.Multiply()=", v.Multiply())
	fmt.Printf("v.X= %v, v.Y=%v\n", v.X, v.Y)
	v.Trans(10)
	fmt.Printf("v.X= %v, v.Y=%v\n", v.X, v.Y)
}
