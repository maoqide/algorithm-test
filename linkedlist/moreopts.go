package linkedlist

// Reverse LinkedList O(n)
func Reverse(l *LinkedList) {
	if l.Head().Next() == nil || l.Head().Next().Next() == nil {
		return
	}
	var cur = l.Head().Next()
	var pre *Node
	for cur != nil {
		tmp := cur.Next()
		cur.SetNext(pre)
		pre = cur
		cur = tmp
	}
	l.head.SetNext(pre)
	return
}

// FindMidNode find mid nodes of the LinkedList
func FindMidNode(l *LinkedList) *Node {
	if l.Head().Next() == nil {
		return nil
	}
	if l.Head().Next().Next() == nil {
		return l.Head().Next()
	}
	var p = l.Head().Next() // 慢指针
	var q = l.Head().Next() // 快指针
	for q.Next() != nil && q.Next().Next() != nil {
		p = p.Next()
		q = q.Next().Next()
	}
	return p
}

// HasCycle judge if LinkedList has cycle
func HasCycle(l *LinkedList) bool {
	var p = l.Head().Next()
	var q = l.Head().Next()
	for q.Next() != nil && q.Next().Next() != nil {
		p = p.Next()
		q = q.Next().Next()
		if p == q {
			return true
		}
	}
	return false
}

// MergeSortedList merge sorted LinkedList, value of Node must be int
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1.Head().Next() == nil {
		return l2
	}
	if l2.Head().Next() == nil {
		return l1
	}
	var result = NewLinkedList()
	var cur = result.Head()
	var cur1 = l1.Head().Next()
	var cur2 = l2.Head().Next()
	for cur1 != nil && cur2 != nil {
		if cur1.Value().(int) <= cur2.Value().(int) {
			cur.SetNext(cur1)
			cur1 = cur1.Next()
		} else {
			cur.SetNext(cur2)
			cur2 = cur2.Next()
		}
		cur = cur.Next()
	}
	if cur1 != nil {
		cur.SetNext(cur1)
	} else if cur2 != nil {
		cur.SetNext(cur2)
	}
	return result
}

// DeleteBackN delete number N Node  count backwards
func DeleteBackN(l *LinkedList, n int) {
	if n <= 0 || l.Head().Next() == nil {
		return
	}

	fast := l.Head().Next()
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next()
	}

	if fast == nil {
		return
	}

	slow := l.Head().Next()
	for fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next()
	}
	slow.SetNext(slow.Next().Next())
	return
}
