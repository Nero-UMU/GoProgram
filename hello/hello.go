package main

import (
	"fmt"

	"example.com/greetings"

	"add"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	seq := add.Sum(10, 20)
	fmt.Println(seq)
}
