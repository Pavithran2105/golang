package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// the first argument i.e. program name is excluded
	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength)
	var z []string
	var x []int
	for i, a := range os.Args[1:] {

		if a[0:1] == "-" {
			fmt.Printf("function %d is %s\n", i+1, a)
			a = a[1:]

			z = append(z, a)
			fmt.Printf("values %d is %s\n", i+1, a)
		} else {
			value, _ := strconv.ParseInt(a, 0, 64)
			x = append(x, int(value))

		}

	}
	fmt.Println("this are the values", x)
	r := x[0]
	u := x[1]
	fmt.Println(r)
	fmt.Println(u)
	fmt.Println(z)
	// res1 := strings.Split(z, ",")
	// fmt.Println(res1)
	n := []string{}
	for _, name := range z {
		// fmt.Println(name)
		res1 := strings.Split(name, "")
		fmt.Println(res1)
		// m := []int{res1}
		n = append(n, res1...)
	}
	fmt.Println(n)
	for _, value := range n {
		// fmt.Println(value)
		switch value {
		case "a":
			fmt.Println(r + u)
		case "s":
			fmt.Println(r - u)
		case "m":
			fmt.Println(r * u)
		case "d":
			fmt.Println(r / u)
		default:
			fmt.Println("this value is not assigned ", value)
		}
	}

}
