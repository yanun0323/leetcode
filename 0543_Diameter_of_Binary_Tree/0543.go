package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	maxDepth(root, &result)
	return result
}

func maxDepth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	lMax := maxDepth(node.Left, diameter)
	rMax := maxDepth(node.Right, diameter)
	if *diameter < lMax+rMax {
		*diameter = lMax + rMax
	}
	return 1 + max(lMax, rMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
