package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	temp1 := make([][]uint8, dy)
	for i := range temp1 {
		temp2 := make([]uint8, dx)
		for j := range temp2 {
			temp2[j] = uint8((i + j) / 2)
		}
		temp1[i] = temp2
	}
	return temp1
}

func main() {
	pic.Show(Pic)
}
