package serverManager

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	commonObjects "reverse_proxy/commons"
	"reverse_proxy/config"
	loadbalancingstrategies "reverse_proxy/utils"
)

type ServerManager interface {
	Initialize()
	HandleRequest()
}

type ReverseProxy struct {
	servers      []commonObjects.BackendServer
	loadBalancer loadbalancingstrategies.LoadBalancer
}

func (r *ReverseProxy) Initialize() {
	for _, su := range config.ServerUrls {
		r.servers = append(r.servers, commonObjects.BackendServer{ServerUrl: su, ActiveRequestCount: 0})
	}
	log.Printf("Servers Initialised....")

	r.loadBalancer = loadbalancingstrategies.GelLoadBalancer()
	r.loadBalancer.Initialize(r.servers)
	log.Printf("Load Balander Initialised....")
}

func (rp *ReverseProxy) HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

	backend := rp.loadBalancer.SelectBackend()
	targetUrl, err := url.Parse(backend.ServerUrl)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("%v", backend.ServerUrl)
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ServeHTTP(w, r)
	rp.loadBalancer.RequestCompleted(backend)

}
