package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	// scanner1 := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintf(conn, "do you want to encrypt a text")
		ln := scanner.Text()
		if ln == "y" {
			fmt.Fprintf(conn, "start")
			l := scanner.Text()

			fmt.Println(l)
			fmt.Fprintf(conn, "I heard you say: %s\n", l)
			// l := scanner.Text()
			// r := []rune(l)

			// fmt.Println(r)
		}
		// fmt.Printf("%U\n", r)
		// fmt.Println(ln)
		// fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	// for scanner1.Scan() {
	// 	lo := scanner1.Text()
	// 	// r1 := []string(lo)
	// 	fmt.Println(lo)
	// }
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
