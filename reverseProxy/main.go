package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
func forwardRequest(res http.ResponseWriter, req *http.Request) {
	url := server()
	rproxy := httputil.NewSingleHostReverseProxy(url)
	rproxy.ServeHTTP(res, req)
}
func server() *url.URL {
	url, _ := url.Parse("http://127.0.0.1:8000")
	return url

}
