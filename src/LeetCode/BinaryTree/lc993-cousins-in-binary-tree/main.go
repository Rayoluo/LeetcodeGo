package main

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/17 9:30
 * @Desc: 二叉树的堂兄弟节点，二叉树的节点值都是唯一的，二叉树的节点数介于2和100之间，root非空
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var parent map[int]*TreeNode

func isCousins(root *TreeNode, x int, y int) bool {
	parent = make(map[int]*TreeNode)
	parent[root.Val] = nil
	findParent(root)
	if parent[x] == parent[y] {
		return false
	}
	// 判断两个节点是否在同一深度
	val := x
	var (
		d1, d2 int
	)
	for parent[val] != nil {
		val = parent[val].Val
		d1++
	}
	val = y
	for parent[val] != nil {
		val = parent[val].Val
		d2++
	}
	if d1 == d2 {
		return true
	}
	return false
}

func findParent(root *TreeNode) {
	if root.Left != nil {
		parent[root.Left.Val] = root
		findParent(root.Left)
	}
	if root.Right != nil {
		parent[root.Right.Val] = root
		findParent(root.Right)
	}
}
