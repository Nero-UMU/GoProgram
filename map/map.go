package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["route"] = 66
	i, ok := m["route"]
	println("i = ", i)
	println("ok = ", ok)

	length := len(m)
	println("length = ", length)

	j := m["root"]
	println("j = ", j)

	comits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	for key, value := range comits {
		fmt.Println("key = ", key, "value = ", value)
	}
}
