package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	loadbalancingstrategies "reverse_proxy/utils"
)

var loadBalancer loadbalancingstrategies.LoadBalancer = loadbalancingstrategies.GelLoadBalancer()

func main() {
	http.HandleFunc("/", handleRequest)

	portAddress := 8080
	fmt.Println("Starting the server on : http://localhost:", portAddress, "!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

	backend := loadBalancer.SelectBackend()
	targetUrl, err := url.Parse(backend.ServerUrl)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("%v", backend.ServerUrl)
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ServeHTTP(w, r)
}
