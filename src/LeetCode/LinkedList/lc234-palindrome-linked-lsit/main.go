package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 第一种方法 时间、空间复杂度为O(n)
func isPalindrome(head *ListNode) bool {
	arr := []int{}
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	p1 := head
	p2 := head
	preMiddle := &ListNode{Next: head}
	for p1 != nil && p2 != nil && p2.Next != nil {
		p1 = p1.Next
		preMiddle = preMiddle.Next
		p2 = p2.Next.Next
	}
	// p1为链表的中间节点
	var pre *ListNode
	cur := p1
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	preMiddle.Next = pre
	cur, p1 = pre, head
	for cur != nil && p1 != pre {
		if cur.Val != p1.Val {
			return false
		}
		cur = cur.Next
		p1 = p1.Next
	}
	return true
}
