package main

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/7/8 0:59
 * @Desc: 相交链表
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		ma, mb *ListNode
	)
	ma, mb = headA, headB
	for ma != mb {
		if ma == nil {
			ma = headB
		} else {
			ma = ma.Next
		}
		if mb == nil {
			mb = headA
		} else {
			mb = mb.Next
		}
	}
	return ma
}
