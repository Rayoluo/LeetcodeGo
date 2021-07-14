package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	root := &TreeNode{
		0,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	fmt.Println(depth(root))
}
