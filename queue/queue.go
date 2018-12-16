package queue

// Queue interface
type Queue interface {
	Enqueue(v interface{})
	DeQueue()
}
