package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	result := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if k == 0 || root == nil {
			return
		}
		dfs(root.Left)
		k--
		if k == 0 {
			result = root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return result
}
