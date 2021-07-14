package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	findParent(root, parent)
	pset := make(map[*TreeNode]struct{})
	for p != nil {
		pset[p] = struct{}{}
		p = parent[p.Val]
	}

	for q != nil {
		if _, exist := pset[q]; exist {
			return q
		}
		q = parent[q.Val]
	}
	return root
}

func findParent(root *TreeNode, parent map[int]*TreeNode) {
	if root.Left != nil {
		parent[root.Left.Val] = root
		findParent(root.Left, parent)
	}
	if root.Right != nil {
		parent[root.Right.Val] = root
		findParent(root.Right, parent)
	}
}
