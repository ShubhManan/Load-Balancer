package main

import (
	"fmt"
	"log"
	"net/http"
	serverManager "reverse_proxy/services"
)

var reverseProxy serverManager.ReverseProxy

func init() {
	reverseProxy.Initialize()
}

func main() {
	http.HandleFunc("/", reverseProxy.HandleRequest)

	portAddress := 8080
	fmt.Println("Starting the server on : http://localhost:", portAddress, "!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
