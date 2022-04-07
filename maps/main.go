package main

import (
	"fmt"
)

type emp struct {
	name string
	ageo int
}

func main() {
	age := make(map[string]int)
	age["pavithran"] = 18
	fmt.Println(age)

	p := emp{
		"pavith",
		21,
	}
	fmt.Println(p)
}
