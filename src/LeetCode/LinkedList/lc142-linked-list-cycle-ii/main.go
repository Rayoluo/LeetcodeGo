package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 判断链表是否有环，如果有环就返回入环的第一个节点
func detectCycle(head *ListNode) *ListNode {
	var (
		fast, slow *ListNode
	)
	fast = head
	slow = head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 此时链表有环
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return slow
		}
	}
	// 链表无环，返回空节点
	return nil
}
