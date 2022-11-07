package main

import (
	"flag"
	"fmt"
)

type Name string

func (i *Name) String() string {
	*i = "Default is me~"
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if value == "hello" {
		*i = "world"
		return nil
	}
	*i = Name("Go tour hajimaliyo~ " + value)
	return nil
}

func main() {
	var name Name
	flag.Var(&name, "name", "help")
	flag.Parse()
	fmt.Println(name)
}
