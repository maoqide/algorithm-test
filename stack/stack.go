package stack

import "fmt"

// Stack interface
type Stack interface {
	Pop() interface{}
	Push(v interface{})
	Top() interface{}
	IsEmpty() bool
	ToString() string
}

// ArrayStack implement Stack based on array
type ArrayStack struct {
	data []interface{}
	top  int
}

// NewArrayStack create a arrayStack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// IsEmpty judge if stack is empty
func (a *ArrayStack) IsEmpty() bool {
	return a.top < 0
}

//Pop pop from stack
func (a *ArrayStack) Pop() interface{} {
	if a.top < 0 {
		return nil
	}
	result := a.data[a.top]
	a.top = a.top - 1
	return result
}

//Push push to stack
func (a *ArrayStack) Push(v interface{}) {
	top := a.top + 1

	if a.top < len(a.data)-1 {
		a.data[top] = v
	} else {
		a.data = append(a.data, v)
	}
	a.top = top
	return
}

// Top of stack
func (a *ArrayStack) Top() interface{} {
	if a.top < 0 {
		return nil
	}
	return a.data[a.top]
}

// ToString convert to string for print
func (a *ArrayStack) ToString() string {
	result := ""
	for i := a.top; i >= 0; i-- {

		result += fmt.Sprintf("%+v", a.data[i])
		if i > 0 {
			result += ","
		}
	}
	return result
}

// LinkedListStack Stack implemented based on LinkedList
type LinkedListStack struct {
	top *node
}

// Node of LinkedList
type node struct {
	value interface{}
	next  *node
}

func newNode(v interface{}) *node {
	return &node{value: v, next: nil}
}

// NewLinkedListStack create a LinkedListStack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

// IsEmpty judge if stack is empty
func (l *LinkedListStack) IsEmpty() bool {
	if l.top == nil {
		return true
	}
	return false
}

//Pop pop from stack
func (l *LinkedListStack) Pop() interface{} {
	top := l.top
	if top != nil {
		l.top = l.top.next
	}
	return top.value
}

//Push push to stack
func (l *LinkedListStack) Push(v interface{}) {
	top := newNode(v)
	top.next = l.top
	l.top = top
	return
}

// Top of the stack
func (l *LinkedListStack) Top() interface{} {
	if l.top == nil {
		return nil
	}
	return l.top.value
}

// ToString convert to string for print
func (l *LinkedListStack) ToString() string {
	result := ""
	cur := l.top
	for cur != nil {
		result += fmt.Sprintf("%+v", cur.value)
		if cur.next != nil {
			result += ","
		}
		cur = cur.next
	}
	return result
}
