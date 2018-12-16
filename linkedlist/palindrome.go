package linkedlist

/*
** 1.快慢指针定位中间节点
** 2.从中间节点对后半部分逆序
** 3.前后半部分比较，判断是否为回文
** 4.后半部分逆序复原
** 时间复杂度O(n), 空间复杂度O(1)
 */

// IsPalindrome judge if LinkedList 'l' palindrome
func IsPalindrome(l *LinkedList) bool {
	if l.Head().Next() == nil {
		return false
	}
	// 节点数为1
	if l.Head().Next().Next() == nil {
		return true
	}
	var isPalindrome = true
	var pre *Node

	var cur = l.Head().Next() // 慢指针
	var q = l.Head().Next()   // 快指针
	for q.Next() != nil && q.Next().Next() != nil {
		q = q.Next().Next()
		tmp := cur.Next()
		cur.SetNext(pre)
		pre = cur
		cur = tmp
	}

	var left, right *Node
	if q.Next() == nil {
		// 奇数节点, cur为中点
		left, right = pre, cur.Next()
	} else {
		// 偶数节点, cur为前一个中点, cur前进一个节点
		tmp := cur.Next()
		cur.SetNext(pre)
		pre = cur
		cur = tmp
		left, right = pre, cur
	}

	for left != nil && right != nil {
		if left.Value() != right.Value() {
			isPalindrome = false
			break
		}
		left = left.Next()
		right = right.Next()
	}

	// 复原链表
	pre, cur = cur, pre
	for cur != nil {
		tmp := cur.Next()
		cur.SetNext(pre)
		pre = cur
		cur = tmp
	}
	l.head.SetNext(pre)
	return isPalindrome

}
