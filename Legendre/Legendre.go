package main

func Legendre(a, p int) int {
	if a%p == 0 {
		return 0
	} else {
		if Euler(a, p) == 1 {
			return 1
		} else {
			return -1
		}
	}
}
