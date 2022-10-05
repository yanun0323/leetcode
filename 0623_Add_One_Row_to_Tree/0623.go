package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}

	if depth > 2 {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
		return root
	}

	if depth == 2 {
		root.Left = &TreeNode{
			Val:  val,
			Left: root.Left,
		}
		root.Right = &TreeNode{
			Val:   val,
			Right: root.Right,
		}
		return root
	}

	return &TreeNode{
		Val:   val,
		Left:  root,
		Right: nil,
	}
}
