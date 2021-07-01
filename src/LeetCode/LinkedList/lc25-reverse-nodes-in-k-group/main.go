package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/7/1 10:53
 * @Desc: k个一组翻转链表
 */

// 仅交换节点值的写法
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		store           []int
		i, j            int
		cur, start, tmp *ListNode
	)
	store = make([]int, k)
	start = head
	for start != nil {
		cur = start
		i = 0
		for cur != nil && i < k {
			store[i] = cur.Val
			cur = cur.Next
			i++
		}

		// 如果当前group有k个元素
		if i == k {
			// 执行翻转操作
			tmp = start
			for j = 0; j < k; j++ {
				tmp.Val = store[k-1-j]
				tmp = tmp.Next
			}
			// 更新start指向下一个k group
			start = cur

		} else { // 如果当前group没有k个元素则交换完成
			break
		}
	}
	return head
}

// k个为一组，反转链表的方式
func reverseKGroup1(head *ListNode, k int) *ListNode {
	var (
		hair, cur, start, pre, prestart *ListNode
		i                               int
	)
	hair = &ListNode{Next: head}
	prestart = hair
	start = head
	for start != nil {
		cur = start
		i = 0
		for cur != nil && i < k {
			pre = cur
			cur = cur.Next
			i++
		}

		// 如果当前group有k个元素
		if i == k {
			// 执行翻转操作
			pre.Next = nil
			tmp := reverseList(start)
			prestart.Next = tmp
			prestart = start
			start.Next = cur
			// 更新start指向下一个k group
			start = cur
		} else { // 如果当前group没有k个元素则交换完成
			break
		}
	}
	return hair.Next
}

// reverseList 反转链表
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
