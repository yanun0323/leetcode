package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if check(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	return root.Val == subRoot.Val && check(root.Left, subRoot.Left) && check(root.Right, subRoot.Right)
}

func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		if queue[0].Val == subRoot.Val && check(queue[0], subRoot) {
			return true
		}
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	return false
}
