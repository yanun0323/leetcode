package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(in []int) *TreeNode
	build = func(in []int) *TreeNode {
		if len(in) == 0 {
			return nil
		}

		i := 0
		for in[i] != preorder[0] {
			i++
		}

		node := &TreeNode{}
		node.Val = preorder[0]
		preorder = preorder[1:]
		node.Left = build(in[:i])
		if i < len(in)-1 {
			node.Right = build(in[i+1:])
		}
		return node
	}
	return build(inorder)
}
