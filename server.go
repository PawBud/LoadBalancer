package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	ReverseProxy *httputil.ReverseProxy
	URL string
	isActive bool
}

func newServer(urlStr string) *server{
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		URL: urlStr,
		ReverseProxy: rp,
		isActive: true,// by Default is it active
	}
}

func (s *server) serverStatus() {
	resp, err := http.Head(s.URL)
	if err != nil{
		log.Println(err)
	}
	if resp.StatusCode != http.StatusOK {
		s.isActive = false
		return
	}
	s.isActive = true
}