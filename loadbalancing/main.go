package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	arrayOfServers = []string{
		"http://127.0.0.1.8000",
		"http://127.0.0.1.8001",
		"http://127.0.0.1.8002",
		"http://127.0.0.1.8003",
		"http://127.0.0.1.8004",
	}
	index = 0
)

func main() {
	http.HandleFunc("/", forwardingServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
func forwardingServer(res http.ResponseWriter, req *http.Request) {
	url := server()

	rproxy := httputil.NewSingleHostReverseProxy(url)
	rproxy.ServeHTTP(res, req)
}
func server() *url.URL {
	nextIndex := (index + 1) % 5
	url, _ := url.Parse(arrayOfServers[index])
	index = nextIndex
	return url

}
