package stack

import (
	"testing"
	// . "github.com/maoqide/algorithm-test/stack"
)

func TestArrayStack(t *testing.T) {
	var as Stack
	as = NewArrayStack()
	t.Log(as.IsEmpty())

	t.Log("as: " + as.ToString())
	for i := 1; i < 11; i++ {
		as.Push(i)
	}
	t.Log("as: " + as.ToString())
	t.Log(as.IsEmpty())
	for i := 0; i < 5; i++ {
		t.Log(as.Pop())
	}
	as.Push(99)
	t.Log("as: " + as.ToString())
}

func TestLinkedListStack(t *testing.T) {
	var ls Stack
	ls = NewLinkedListStack()
	t.Log(ls.IsEmpty())

	t.Log("ls: " + ls.ToString())
	for i := 1; i < 11; i++ {
		ls.Push(i)
	}
	t.Log("ls: " + ls.ToString())
	t.Log(ls.IsEmpty())
	for i := 0; i < 5; i++ {
		t.Log(ls.Pop())
	}
	ls.Push(99)
	t.Log("as: " + ls.ToString())
}
