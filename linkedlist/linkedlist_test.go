package linkedlist

import (
	"testing"
)

func TestBasic(t *testing.T) {
	var data []interface{}
	for i := 0; i < 5; i++ {
		data = append(data, i)
	}
	var l = NewLinkedList()
	l.Dump(data)

	t.Logf("length: %d \n", l.Length())
	t.Log(l.ToString())
	t.Log(l.Find(2))
	l.Insert(l.Head().Next(), 5)
	t.Log(l.ToString())
	l.Delete(l.Head().Next().Next())
	t.Log(l.ToString())
}

func TestIsPalindrome(t *testing.T) {
	var data []interface{}
	data = append(data, []interface{}{1, 2, 3, 4, 4, 3, 2, 1}...)
	var l = NewLinkedList()
	l.Dump(data)

	t.Log(l.ToString())
	t.Log(IsPalindrome(l))
	t.Log(l.ToString())

}

func TestReverse(t *testing.T) {
	var data []interface{}
	data = append(data, []interface{}{1, 2, 3, 4, 5, 6, 7, 8}...)
	var l = NewLinkedList()
	l.Dump(data)

	t.Log(l.ToString())
	Reverse(l)
	t.Log(l.ToString())
}

func TestFindMidNode(t *testing.T) {
	var data []interface{}
	data = append(data, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	var l = NewLinkedList()
	l.Dump(data)

	t.Log(FindMidNode(l))
}

func TestHasCycle(t *testing.T) {
	var data []interface{}
	data = append(data, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	var l = NewLinkedList()
	l.Dump(data)

	t.Logf(l.ToString())
	t.Log(HasCycle(l))
	l.Find(5).SetNext(l.Find(3))
	t.Log(HasCycle(l))
}

func TestMergeSortedList(t *testing.T) {
	var data []interface{}
	data = append(data, []interface{}{1, 2, 5, 9, 21, 23, 33}...)
	var l1 = NewLinkedList()
	l1.Dump(data)
	var data2 []interface{}
	data2 = append(data2, []interface{}{3, 4, 5, 6, 7, 9}...)
	var l2 = NewLinkedList()
	l2.Dump(data2)

	t.Log(MergeSortedList(l1, l2).ToString())
}

func TestDeleteBackN(t *testing.T) {
	var data []interface{}
	data = append(data, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	var l = NewLinkedList()
	l.Dump(data)

	t.Logf(l.ToString())
	DeleteBackN(l, 3)
	t.Logf(l.ToString())
}
