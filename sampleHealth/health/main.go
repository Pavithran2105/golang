package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	serverList = []*server{
		newServer("server-1", "http://127.0.0.1:5001"),
		newServer("server-2", "http://127.0.0.1:5002"),
		newServer("server-3", "http://127.0.0.1:5003"),
		newServer("server-4", "http://127.0.0.1:5004"),
		newServer("server-5", "http://127.0.0.1:5005"),
	}
	lastServedIndex = 0
)

type server struct {
	Name         string
	Url          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func newServer(name, urlStr string) *server {
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		Name:         name,
		Url:          urlStr,
		ReverseProxy: rp,
		Health:       true,
	}
}

func startHealthCheck() {
	s := gocron.NewScheduler(time.Local)
	for _, host := range serverList {
		_, err := s.Every(2).Seconds().Do(func(so *server) {
			healthy := so.checkHealth()
			if healthy {
				log.Printf("'%s' is healthy!", so.Name)
			} else {
				log.Printf("'%s' is not healthy", so.Name)
			}
		}, host)
		if err != nil {
			log.Fatalln(err)
		}
	}
	s.StartAsync()
}
func (si *server) checkHealth() bool {
	resp, err := http.Head(si.Url)
	if err != nil {
		si.Health = false
		return si.Health
	}
	if resp.StatusCode != http.StatusOK {
		si.Health = false
		return si.Health
	}
	si.Health = true
	return si.Health
}

////

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server, err := getHealthyServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(res, req)
}
func getHealthyServer() (*server, error) {
	for i := 0; i < len(serverList); i++ {
		server := getServer()
		if server.Health {
			return server, nil
		}
	}
	return nil, fmt.Errorf("No healthy hosts")
}
func getServer() *server {
	nextIndex := (lastServedIndex + 1) % len(serverList)
	server1 := serverList[nextIndex]
	lastServedIndex = nextIndex
	return server1
}

func main() {
	http.HandleFunc("/", forwardRequest)
	// go startHealthCheck()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
