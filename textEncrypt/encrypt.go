package main

import (
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
		fmt.Println("Do you want to encrypt or decrypt your text?")
		//var answer string
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
		} else if input == "2" {
			for {
				conn, err := li.Accept()
				if err != nil {
					fmt.Println(err)
					return
				}
				go decryptConnection(conn)
			}
		}
	}
}
func decryptConnection(conn net.Conn) {
	fmt.Println("hello")
}
func handleConnection(conn net.Conn) {
	fmt.Println("hello2")
}
