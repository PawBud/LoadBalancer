package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",forwardRequest)
}

func forwardRequest(res http.ResponseWriter, req *http.Request){
	log.Fatal(http.ListenAndServe(":8888", nil))
}
