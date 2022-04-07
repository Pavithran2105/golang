package main

import (
	"fmt"
	"net/http"
)

var (
	strArray1 = [3]string{"Japan", "Australia", "Germany"}
	lastIndex = 0
)

func main() {

	// fmt.Println(strArray1[1])
	http.HandleFunc("/", simpleServer)
	http.ListenAndServe(":8000", nil)
	// fmt.Println(server())
}
func simpleServer(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, server())
}

func server() string {
	nextlastIndex := (lastIndex + 1) % 5
	// fmt.Println(strArray1[lastIndex])
	yes := (strArray1[lastIndex])
	lastIndex = nextlastIndex
	return yes

}
