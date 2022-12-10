package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const mod = 1_000_000_000 + 7

func maxProduct(root *TreeNode) int {
	candidate := []int{}

	var getSum func(n *TreeNode) int
	getSum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := getSum(n.Left)
		r := getSum(n.Right)
		candidate = append(candidate, l, r)
		return l + r + n.Val
	}
	sum := getSum(root)
	gap := sum

	result := 0
	for i := range candidate {
		if g := abs(sum - 2*candidate[i]); g < gap {
			result = candidate[i]
			gap = g
		}
	}

	return ((sum - result) * result) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
