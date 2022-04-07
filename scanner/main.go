package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// fmt.Println("Enter text: ")
	// text2 := ""
	// fmt.Scanln(&text2)
	// fmt.Println(text2)
	// a := 0
	// var reader string
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// 	reader += (scanner.Text())

	// 	if a == 5 {
	// 		fmt.Println(reader)
	// 		break
	// 	}
	// 	a += 1
	// }
	// ln := ""
	// fmt.Sscanln("%v", ln)
	// fmt.Println(ln)
}
