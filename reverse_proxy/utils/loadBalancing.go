package loadbalancingstrategies

import (
	"container/heap"
	"log"
	"os"
	commonObjects "reverse_proxy/commons"
	"reverse_proxy/queue"
)

type LoadBalancer interface {
	Initialize([]commonObjects.BackendServer)
	SelectBackend() *commonObjects.BackendServer
	RequestCompleted(*commonObjects.BackendServer)
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

func (s *RoundRobinLoadBalancer) SelectBackend() *commonObjects.BackendServer {
	svr, _ := s.ServersList.Pop()
	s.ServersList.Push(svr)
	return &svr
}

func (s *RoundRobinLoadBalancer) RequestCompleted(*commonObjects.BackendServer) {}

type LeastConnectionLoadBalancer struct {
	ServersList queue.PriorityQueue
}

func (lc *LeastConnectionLoadBalancer) Initialize(serverList []commonObjects.BackendServer) {
	if len(serverList) == 0 {
		log.Fatal("No backend server found!")
	}
	lc.ServersList = make(queue.PriorityQueue, len(serverList))

	i := 0
	for _, backendUrl := range serverList {
		lc.ServersList[i] = &commonObjects.BackendServer{ServerUrl: backendUrl.ServerUrl, ActiveRequestCount: 0, Index: i}
		i++
	}
	heap.Init(&lc.ServersList)
}

func (lc *LeastConnectionLoadBalancer) SelectBackend() *commonObjects.BackendServer {
	var svr *commonObjects.BackendServer = heap.Pop(&lc.ServersList).(*commonObjects.BackendServer)
	svr.ActiveRequestCount += 1
	heap.Push(&lc.ServersList, svr)

	return svr
}

func (lc *LeastConnectionLoadBalancer) RequestCompleted(svr *commonObjects.BackendServer) {
	svr.ActiveRequestCount -= 1
	heap.Fix(&lc.ServersList, svr.Index)
}

func GelLoadBalancer() LoadBalancer {
	strategy := os.Getenv("LOAD_BALANCING_ALGO")
	var lb LoadBalancer
	switch strategy {
	case "round_robin":
		var defQ queue.Queue[commonObjects.BackendServer]
		lb = &RoundRobinLoadBalancer{defQ}
	default:
		var defQ queue.PriorityQueue
		lb = &LeastConnectionLoadBalancer{defQ}
	}

	return lb
}
