package main

import "fmt"

type Vertex struct {
	X string
}

type MyInter interface {
	fun()
}

type MyFloat float64

func (V *Vertex) fun() {
	fmt.Printf("v.X=%v", V.X)
}

func main() {
	var i interface{} = Vertex{"hello"}

	t, ok := i.(Vertex)
	fmt.Printf("t=%v, ok=%v\n", t, ok)
	s, ok := i.(string)
	fmt.Printf("s=%v, ok=%v\n", s, ok)
	getType(i)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
