package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for layer := 0; len(queue) > 0; layer++ {
		result = append(result, queue[0].Val)
		for i := len(queue); i > 0; i-- {
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			queue = queue[1:]
		}
	}
	return result
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	var recursive func(root *TreeNode, level int)
	recursive = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level == len(result) {
			result = append(result, root.Val)
		}

		if root.Right != nil {
			recursive(root.Right, level+1)
		}

		if root.Left != nil {
			recursive(root.Left, level+1)
		}
	}
	recursive(root, 0)
	return result
}
