package main

import (
	"fmt"
	"time"
)

func say(word string) {
	for i := 0; i < 10; i++ {
		fmt.Println(word)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go say("Hello")
	say("World")
}
