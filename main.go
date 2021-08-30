package main

import (
	"log"
	"net/http"
)

var (
	serverList = []*server{
		newServer("http://127.0.0.1:5001"),
		newServer("http://127.0.0.1:5002"),
		newServer("http://127.0.0.1:5003"),
		newServer("http://127.0.0.1:5004"),
		newServer("http://127.0.0.1:5005"),
	}
	lastServedIndex = 0
)

func main()  {
	http.HandleFunc("/",forwardRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request){
	server := getServer()
	server.ReverseProxy.ServeHTTP(res, req)
}

func getServer() *server {
	nextIndex := (lastServedIndex+1)%len(serverList)
	server := serverList[lastServedIndex]
	lastServedIndex = nextIndex
	return server
}
