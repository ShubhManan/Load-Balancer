package loadbalancingstrategies

import (
	"log"
	"os"
	commonObjects "reverse_proxy/commons"
	"reverse_proxy/queue"
)

type LoadBalancer interface {
	Initialize([]commonObjects.BackendServer)
	SelectBackend() commonObjects.BackendServer
}

type RoundRobinLoadBalancer struct {
	ServersList queue.Queue[commonObjects.BackendServer]
}

func (s *RoundRobinLoadBalancer) Initialize(serverList []commonObjects.BackendServer) {
	if len(serverList) == 0 {
		log.Fatal("No backend server found!")
	}
	for _, backendUrl := range serverList {
		backend := commonObjects.BackendServer{ServerUrl: backendUrl.ServerUrl, ActiveRequestCount: 0}
		s.ServersList.Push(backend)
	}
}

func (s *RoundRobinLoadBalancer) SelectBackend() commonObjects.BackendServer {
	svr, _ := s.ServersList.Pop()
	s.ServersList.Push(svr)
	return svr
}

func GelLoadBalancer() LoadBalancer {
	strategy := os.Getenv("LOAD_BALANCING_ALGO")
	var lb LoadBalancer
	switch strategy {
	case "round_robin":
		var defQ queue.Queue[commonObjects.BackendServer]
		lb = &RoundRobinLoadBalancer{defQ}
	default:
		var defQ queue.Queue[commonObjects.BackendServer]
		lb = &RoundRobinLoadBalancer{defQ}
	}

	return lb
}
