package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	tmp := prev
	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next
		prev.Next, a.Next, b.Next = b, b.Next, a
		prev = a
	}
	return tmp.Next
}
