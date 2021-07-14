package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 使用队列存放层序遍历的结果，用变量count表示当前行元素个数
	var (
		q           []*TreeNode
		cnt, i, tmp int
		node        *TreeNode
		ret         [][]int
	)
	q = []*TreeNode{root}
	ret = make([][]int, 0)
	cnt = 1
	for len(q) != 0 {
		larr := make([]int, 0)
		tmp = 0
		for i = cnt; i > 0; i-- {
			node = q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
				tmp++
			}
			if node.Right != nil {
				q = append(q, node.Right)
				tmp++
			}
			larr = append(larr, node.Val)
		}
		cnt = tmp
		ret = append(ret, larr)
	}
	return ret
}
