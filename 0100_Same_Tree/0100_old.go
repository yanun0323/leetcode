package main

type TreeNode_Old struct {
	Val   int
	Left  *TreeNode_Old
	Right *TreeNode_Old
}

func isSameTree_Old(p *TreeNode_Old, q *TreeNode_Old) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree_Old(p.Left, q.Left) && isSameTree_Old(p.Right, q.Right)
}
