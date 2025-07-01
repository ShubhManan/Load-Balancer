package commons

type BackendServer struct {
	ServerUrl          string
	ActiveRequestCount int
	Index              int // The index of the item in the heap.
}