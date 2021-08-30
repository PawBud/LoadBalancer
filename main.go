package main

import (
	"fmt"
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
	server, err := getActiveServer()
	if err != nil {
		fmt.Fprintf(res, "Couldn't process the request: %s", err.Error())
	}
	server.ReverseProxy.ServeHTTP(res, req)
}

func getActiveServer() (*server, error){
	for i:=0; i<len(serverList); i++ {
		server := getServer()
		if server.isActive {
			return server, nil
		}
	}
	return nil, fmt.Errorf("No Active Servers Remaining!!")
}

func getServer() *server {
	nextIndex := (lastServedIndex+1)%len(serverList)
	server := serverList[lastServedIndex]
	lastServedIndex = nextIndex
	return server
}
