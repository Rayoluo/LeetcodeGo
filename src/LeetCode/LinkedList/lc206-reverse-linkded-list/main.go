package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/6/15 9:30
 * @Desc: 给定链表的头节点，反转链表
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	if head.Next != nil {
// 		ret := reverseList(head.Next)
// 		head.Next.Next = head
// 		return ret
// 	}
// 	return head
// }

// 迭代的方式
func reverseList(head *ListNode) *ListNode {
	var (
		pre, cur, next *ListNode
	)
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: 5, Next: nil}
	ret := reverseList(head)
	fmt.Println(ret.Val)
	fmt.Println(ret.Next.Val)
	fmt.Println(ret.Next.Next.Val)
}
