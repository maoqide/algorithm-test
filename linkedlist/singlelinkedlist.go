package linkedlist

import (
	"fmt"
)

// Node of SingleLinkedList
type Node struct {
	value interface{}
	next  *Node
}

// NewNode create a new Node
func NewNode(v interface{}) *Node {
	return &Node{value: v, next: nil}
}

// Next of the Node 'n'
func (n *Node) Next() *Node {
	return n.next
}

// Value of Node 'n'
func (n *Node) Value() interface{} {
	return n.value
}

// SetNext set Node n.next to 'nxt'
func (n *Node) SetNext(nxt *Node) {
	n.next = nxt
	return
}

// LinkedList is singleLenkedList
type LinkedList struct {
	head *Node // head为哨兵，无意义
}

// NewLinkedList create a new LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{head: NewNode(0)}
}

// Length of LinkedList 'l'
func (l *LinkedList) Length() int {
	var length = 0
	var cur = l.Head().Next()
	for cur != nil {
		length++
		cur = cur.Next()
	}
	return length
}

// Head of LinkedList
func (l *LinkedList) Head() *Node {
	return l.head
}

// Insert a Node into LinkedList 'l' with value 'v' after Node 'n'.
func (l *LinkedList) Insert(n *Node, v interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewNode(v)
	originNext := n.next
	n.next = newNode
	newNode.next = originNext
	return true
}

// Delete Node 'n' from LinkedList 'l'
func (l *LinkedList) Delete(n *Node) bool {
	if n == nil {
		return false
	}
	var cur = l.head.next
	var pre = cur
	for cur != nil {
		if cur == n {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		// n not found
		return false
	}
	pre.next = n.next
	n = nil
	return true
}

// Find Node in LinkedList 'l' with value=v
func (l *LinkedList) Find(v interface{}) *Node {
	if v == nil {
		return nil
	}
	var cur = l.head.next
	for cur != nil {
		if cur.value == v {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// ToString convert LinkedList 'l' to string
func (l *LinkedList) ToString() string {
	var cur = l.head.next
	var result = ""
	for cur != nil {
		result += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if cur != nil {
			result += "->"
		}
	}
	return result
}

// Dump data from array to LinkedList, this will overwirte origin data in LinkedList
func (l *LinkedList) Dump(d []interface{}) {
	var cur = l.head
	for _, v := range d {
		cur.next = NewNode(v)
		cur = cur.next
	}
	return
}
