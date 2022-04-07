package main

import (
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
)

type server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func newServer(name, urlStr string) *server {
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		Name:         name,
		URL:          urlStr,
		ReverseProxy: rp,
		Health:       true,
	}
}
func (s *server) checkHealth() bool {
	resp, err := http.Head(s.URL)
	if err != nil {
		s.Health = false
		return s.Health
	}
	if resp.StatusCode != http.StatusOK {
		s.Health = false
		return s.Health
	}
	s.Health = true
	return s.Health
}

func startHealthCheck() {
	s := gocron.NewScheduler(time.Local)
	for _, host := range serverList {
		_, err := s.Every(2).Seconds().Do(func(s *server) {
			healthy := s.checkHealth()
			if healthy {
				log.Printf("'%s' is healthy!", s.Name)
			} else {
				log.Printf("'%s' is not healthy", s.Name)
			}
		}, host)
		if err != nil {
			log.Fatalln(err)
		}
	}
	s.StartAsync()
}
func main() {

	startHealthCheck()

}
