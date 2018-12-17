package queue

import (
	"testing"
)

func TestArrayQueue(t *testing.T) {
	var q Queue
	q = NewArrayQueue()
	t.Log("q: ", q.ToString(), "len: ", q.Length())

	for i := 0; i < 17; i++ {
		if !q.EnQueue(i) {
			t.Log("queue full")
			break
		}
	}
	t.Log("q: ", q.ToString())
	for i := 0; i < 3; i++ {
		t.Log(q.DeQueue())
	}
	for i := 0; i < 3; i++ {
		t.Log(q.EnQueue(2333))
		t.Log("q: ", q.ToString(), "len: ", q.Length())
	}

}

// func TestLinkedListQueue(t *testing.T) {
// 	var q Queue
// 	q = NewLinkedlistQueue()
// 	t.Log("q: ", q.ToString())

// 	for i := 0; i < 17; i++ {
// 		if !q.EnQueue(i) {
// 			t.Log("queue full")
// 			break
// 		}
// 	}
// 	t.Log("q: ", q.ToString())
// 	for i := 0; i < 1; i++ {
// 		t.Log(q.DeQueue())
// 	}
// 	t.Log("q: ", q.ToString())
// 	t.Log(q.EnQueue(2333))
// 	t.Log("q: ", q.ToString())
// 	t.Log(q.EnQueue(2333))
// 	t.Log("q: ", q.ToString())

// }

func TestCircleQueue(t *testing.T) {
	var q Queue
	q = NewCircleQueue()
	t.Log("q: ", q.ToString(), "len: ", q.Length())

	for i := 0; i < 17; i++ {
		if !q.EnQueue(i) {
			t.Log("queue full")
			break
		}
	}
	t.Log("q: ", q.ToString())
	for i := 0; i < 3; i++ {
		t.Log(q.DeQueue())
	}
	for i := 0; i < 3; i++ {
		t.Log(q.EnQueue(2333))
		t.Log("q: ", q.ToString(), "len: ", q.Length())
	}

}
