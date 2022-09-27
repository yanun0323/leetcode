package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) (count int) {
	countGood(root, root.Val, &count)
	return
}

func countGood(node *TreeNode, maxVal int, count *int) {
	if node == nil {
		return
	}
	if node.Val >= maxVal {
		*count++
	}
	maxVal = max(node.Val, maxVal)
	countGood(node.Left, maxVal, count)
	countGood(node.Right, maxVal, count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
