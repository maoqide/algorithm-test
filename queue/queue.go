package queue

import (
	"fmt"
)

// Queue interface
type Queue interface {
	EnQueue(v interface{}) bool
	DeQueue() interface{}
	ToString() string
	Length() int
}

// ArrayQueue based on array
type ArrayQueue struct {
	capacity int
	data     []interface{}
	head     int
	tail     int
}

// NewArrayQueue create array queue
func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		capacity: 16,
		data:     make([]interface{}, 16),
		head:     0,
		tail:     0,
	}
}

// NewArrayQueueWithCap create array queue
func NewArrayQueueWithCap(cap int) *ArrayQueue {
	return &ArrayQueue{
		capacity: cap,
		data:     make([]interface{}, cap),
		head:     0,
		tail:     0,
	}
}

// DeQueue from queue
func (q *ArrayQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	r := q.data[q.head]
	q.head++
	return r
}

// EnQueue1 v to queue, without data moving
func (q *ArrayQueue) EnQueue1(v interface{}) bool {
	if q.tail == q.capacity {
		return false
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

// EnQueue v to queue
func (q *ArrayQueue) EnQueue(v interface{}) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		q.move()
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) move() {
	fmt.Println("ArrayQueue moving data...")
	for i := 0; i < q.tail-q.head; i++ {
		q.data[i] = q.data[q.head+i]
	}
	q.tail = q.tail - q.head
	q.head = 0
}

// Length of queue
func (q *ArrayQueue) Length() int {
	return q.tail - q.head
}

// ToString convert to string
func (q *ArrayQueue) ToString() string {
	result := ""
	for i := q.head; i < q.tail; i++ {
		result += fmt.Sprintf("%v", q.data[i])
		if i < q.tail-1 {
			result += ", "
		}
	}
	return result
}

// LinkedListQueue based on linkedlist
type LinkedListQueue struct {
	head   *node
	tail   *node
	length int
}

// Node of LinkedList
type node struct {
	value interface{}
	next  *node
}

func newNode(v interface{}) *node {
	return &node{value: v, next: nil}
}

// NewLinkedlistQueue create linkedlist queue
func NewLinkedlistQueue() *LinkedListQueue {
	return &LinkedListQueue{head: nil, tail: nil, length: 0}
}

// DeQueue from queue
func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	r := q.head.value
	q.head = q.head.next
	q.length--
	return r
}

// EnQueue v to queue
func (q *LinkedListQueue) EnQueue(v interface{}) bool {
	n := newNode(v)
	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.length++
	return true
}

// ToString convert to string
func (q *LinkedListQueue) ToString() string {
	result := ""
	cur := q.head
	for cur != nil {
		result += fmt.Sprintf("%v", cur.value)
		if cur.next != nil {
			result += ", "
		}
		cur = cur.next
	}
	return result
}

// Length of queue
func (q *LinkedListQueue) Length() int {
	return q.length
}

// CircleQueue based on array
type CircleQueue struct {
	capacity int
	data     []interface{}
	head     int
	tail     int
}

// NewCircleQueue create circle queue
func NewCircleQueue() *CircleQueue {
	return &CircleQueue{
		capacity: 16 + 1,
		data:     make([]interface{}, 16+1),
		head:     0,
		tail:     0,
	}
}

// NewCircleQueueWithCap create circle queue with custom capacity
func NewCircleQueueWithCap(cap int) *CircleQueue {
	return &CircleQueue{
		capacity: cap + 1,
		data:     make([]interface{}, cap+1),
		head:     0,
		tail:     0,
	}
}

// DeQueue from queue
func (q *CircleQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	r := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	return r
}

// EnQueue v to queue
func (q *CircleQueue) EnQueue(v interface{}) bool {
	if q.full() {
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *CircleQueue) full() bool {
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

// Length of queue
func (q *CircleQueue) Length() int {
	return (q.tail - q.head + q.capacity) % q.capacity
}

// ToString convert to string
func (q *CircleQueue) ToString() string {
	result := ""
	i := q.head
	for i != q.tail {
		result += fmt.Sprintf("%v", q.data[i])
		if i != q.tail {
			result += ", "
		}
		i = (i + 1) % q.capacity
	}
	return result
}
