package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == p.Val || node.Val == q.Val {
			return node
		}

		l := dfs(node.Left)
		r := dfs(node.Right)
		if l != nil && r != nil {
			return node
		}

		if l != nil {
			return l
		}

		if r != nil {
			return r
		}

		return nil
	}
	return dfs(root)
}
