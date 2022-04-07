package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("happy2.txt")
	err = os.Rename("happy2.txt", "happy3.txt")
	data, err := ioutil.ReadFile("happy3.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(file)
	fmt.Print(string(data))
}
