package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [3,null,30,10,null,null,15,null,45]
//       3
//     X   30
//        10  X
//       X 15
//        X 45
func isValidBST(root *TreeNode) bool {
	return DFS(root.Left, nil, &root.Val) && DFS(root.Right, &root.Val, nil)
}

func DFS(node *TreeNode, minVal, maxVal *int) bool {
	if node == nil {
		return true
	}

	if minVal != nil && node.Val <= *minVal {
		return false
	}

	if maxVal != nil && node.Val >= *maxVal {
		return false
	}

	return DFS(node.Left, minVal, &node.Val) && DFS(node.Right, &node.Val, maxVal)
}
