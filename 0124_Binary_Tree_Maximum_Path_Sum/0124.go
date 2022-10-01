package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const ZERO = 0

func maxPathSum(root *TreeNode) int {
	result := root.Val
	var maxPath func(node *TreeNode) int
	maxPath = func(node *TreeNode) int {
		if node == nil {
			return ZERO
		}
		l, r := maxPath(node.Left), maxPath(node.Right)
		n := node.Val + max(l, ZERO) + max(r, ZERO)
		if n > result {
			result = n
		}
		return node.Val + max(max(l, r), ZERO)
	}
	maxPath(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
