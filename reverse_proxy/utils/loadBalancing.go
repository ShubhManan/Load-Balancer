package loadbalancingstrategies

import (
	"log"
	"os"
	"reverse_proxy/config"
	"reverse_proxy/queue"
)

type BackendServer struct {
	ServerUrl string
}

type LoadBalancer interface {
	Initialize([]BackendServer)
	SelectBackend() BackendServer
}

type RoundRobinLoadBalancer struct {
	ServersList queue.Queue[BackendServer]
}

func (s *RoundRobinLoadBalancer) Initialize(serverList []BackendServer) {
	if len(serverList) == 0 {
		log.Fatal("No backend server found!")
	}
	for _, backendUrl := range serverList {
		backend := BackendServer{backendUrl.ServerUrl}
		s.ServersList.Push(backend)
	}
}

func (s *RoundRobinLoadBalancer) SelectBackend() BackendServer {
	svr, _ := s.ServersList.Pop()
	s.ServersList.Push(svr)
	return svr
}

func GelLoadBalancer() LoadBalancer {
	strategy := os.Getenv("LOAD_BALANCING_ALGO")
	var lb LoadBalancer
	switch strategy {
	case "round_robin":
		var defQ queue.Queue[BackendServer]
		lb = &RoundRobinLoadBalancer{defQ}
	default:
		var defQ queue.Queue[BackendServer]
		lb = &RoundRobinLoadBalancer{defQ}
	}

	// Initialize the servers and the load balancing strategy
	var servers []BackendServer
	for _, su := range config.ServerUrls {
		servers = append(servers, BackendServer{su})
	}
	lb.Initialize(servers)

	return lb
}
