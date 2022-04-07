package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	takingArgs = os.Args[1:]
)

func main() {

	http.HandleFunc("/", simpleServer)
	log.Fatal(http.ListenAndServe(takingArgs[0], nil))
}

func simpleServer(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "this is server %s", takingArgs[0])
	//fmt.Fprintf(res, "hi")
}

// a='/home/spurge/goNetworks/simpleServerArgs'
// cd $a;

// for i in {1..5}
// do
//    go run main.go ":800$i"
// done
