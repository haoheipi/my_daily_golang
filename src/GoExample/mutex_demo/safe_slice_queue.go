package main

import "sync"

type SafeSliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SafeSliceQueue) {
	return &SafeSliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SafeSliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SafeSliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}

func main() {
	queue := NewSliceQueue(10)
	queue.Enqueue("a")
	queue.Enqueue(1)
	queue.Dequeue()
	queue.Dequeue()
}
