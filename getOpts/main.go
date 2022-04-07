package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength)
	//it gives the number of arguments passed
	//it takes both -flag and values
	var argv []string
	var arg []int
	for _, a := range os.Args[1:] {

		if a[0:1] == "-" {
			//this if condition will check if the arguments start with "-" i.e flag
			// fmt.Printf("function %d is %s\n", i+1, a)
			a = a[1:]
			//this will slice out the "-" flag
			argv = append(argv, a)
			// fmt.Printf("values %d is %s\n", i+1, a)
		} else {
			value, _ := strconv.ParseInt(a, 0, 64)
			//string value is converted to int value
			arg = append(arg, int(value))
			//if the argument doesnt start with flag then it is a value
		}

	}
	// fmt.Println("this are the values", arg)
	firstArgumentValue := arg[0]
	secondArgumentValue := arg[1]
	// fmt.Println(firstArgumentValue)
	// fmt.Println(secondArgumentValue)
	// fmt.Println(argv)
	// res1 := strings.Split(z, ",")
	// fmt.Println(res1)
	casesStoredInArrays := []string{}
	for _, name := range argv {
		// fmt.Println(name)
		res1 := strings.Split(name, "")
		// fmt.Println(res1)
		// m := []int{res1}
		casesStoredInArrays = append(casesStoredInArrays, res1...)
	}
	fmt.Println("these are the funtions you choosed", casesStoredInArrays)
	for _, getOPts := range casesStoredInArrays {
		// fmt.Println(getOPts)
		switch getOPts {
		case "a":
			fmt.Println("addition", firstArgumentValue+secondArgumentValue)
		case "s":
			fmt.Println("subraction", firstArgumentValue-secondArgumentValue)
		case "m":
			fmt.Println("multiplication", firstArgumentValue*secondArgumentValue)
		case "d":
			fmt.Println("division", firstArgumentValue/secondArgumentValue)
		default:
			fmt.Println("this function is not assigned ", getOPts)
		}
	}

}

//go build main.go
// below are the test cases you can run
// ./main -asmd 55 46
// ./main 78 90 -a
// ./main -a 60 70 -sm
