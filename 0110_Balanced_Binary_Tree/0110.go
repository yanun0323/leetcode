package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := maxDepth(root)
	return isBalanced
}

func maxDepth(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	lMax, lIsBalanced := maxDepth(node.Left)
	rMax, rIsBalanced := maxDepth(node.Right)
	if !lIsBalanced || !rIsBalanced {
		return 0, false
	}

	gap := lMax - rMax
	if gap > 1 || gap < -1 {
		return 0, false
	}

	return 1 + max(lMax, rMax), true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
