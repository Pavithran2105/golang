package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(mini())

}
func mini() string {
	u := "hi"
	t := "hello"
	e := flag.String("u", "", "enter username")
	flag.Parse()
	z := u + " " + *e + t

	return z
}
