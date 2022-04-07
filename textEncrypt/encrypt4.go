package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		fmt.Println("Do you want to encrypt or decrypt your text?")
		// var answer string
		var input string
		_, err = fmt.Scanln(&input)
		if input == "1" {
			for {
				conn, err := li.Accept()
				if err != nil {
					fmt.Println(err)
					return
				}
				go handleConnection(conn)
			}

			// } else if input == "2" {
			// 	for {
			// 		conn, err := li.Accept()
			// 		if err != nil {
			// 			fmt.Println(err)
			// 			return
			// 		}
			// 		go decryptConnection(conn)
			// 	}
		}
	}

}
func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		r := []rune(ln)
		fmt.Println(r)

		fmt.Fprint(conn, r)

	}
	defer conn.Close()

	fmt.Println("Code got here.")
}

// func decryptConnection(conn net.Conn) {
// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {

// 		ln := scanner.Text()
// 		c := []rune(ln)
// 		// fmt.Fprint(r)
// 		// ra := r.Next()
// 		// fmt.Fprint(conn, strconv.QuoteRune(ra))
// 		// c := ln.Next()
// 		fmt.Fprint(conn, strconv.QuoteRune(c))
// 	}
// 	defer conn.Close()

// 	fmt.Println("Code got here.")
// }
