package queue

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes an item from the queue
func (q *Queue[T]) Pop() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}
