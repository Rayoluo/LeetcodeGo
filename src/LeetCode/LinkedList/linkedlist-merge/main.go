package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func mergeSortList(headA, headB *Node) *Node {
	if headA == nil {
		return headB
	}
	if headB == nil {
		return headA
	}
	if headA.Val < headB.Val {
		headA.Next = mergeSortList(headA.Next, headB)
		return headA
	} else {
		headB.Next = mergeSortList(headA, headB.Next)
		return headB
	}
}

func main() {
	list1 := &Node{1, &Node{2, nil}}
	list2 := &Node{0, &Node{4, nil}}
	mergeList := mergeSortList(list1, list2)
	fmt.Println(mergeList.Val)
	fmt.Println(mergeList.Next.Val)
	fmt.Println(mergeList.Next.Next.Val)
	fmt.Println(mergeList.Next.Next.Next.Val)
}
