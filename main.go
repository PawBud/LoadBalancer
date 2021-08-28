package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main()  {
	http.HandleFunc("/",forwardRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request){
	url := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(res, req)
}

func getServer() *url.URL {
	url, _ := url.Parse("http://127.0.0.1:4000")
	return url
}
