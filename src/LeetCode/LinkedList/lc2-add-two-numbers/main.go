package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := l1
	b := l2
	prev := &ListNode{Val: -1}
	prehead := prev
	cf := 0
	for a != nil && b != nil {
		res := a.Val + b.Val + cf
		cf = res / 10
		prev.Next = &ListNode{Val: res % 10}
		prev = prev.Next
		a, b = a.Next, b.Next
	}
	for a != nil {
		res := a.Val + cf
		cf = res / 10
		prev.Next = &ListNode{Val: res % 10}
		prev = prev.Next
		a = a.Next
	}
	for b != nil {
		res := b.Val + cf
		cf = res / 10
		prev.Next = &ListNode{Val: res % 10}
		prev = prev.Next
		b = b.Next
	}
	if cf != 0 {
		prev.Next = &ListNode{Val: cf}
	}
	return prehead.Next
}
