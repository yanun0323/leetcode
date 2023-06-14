package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	_max = 1<<(32-1) - 1
)

func getMinimumDifference(root *TreeNode) int {
	result, _, _ := dps(root)
	return result
}

func dps(n *TreeNode) (dif, max, min int) {
	dif, max, min = _max, n.Val, n.Val
	if n.Left != nil {
		lDif, lMax, lMin := dps(n.Left)
		dif = mins(dif, lDif)
		dif = mins(dif, abs(lMax-n.Val))
		min = lMin
	}

	if n.Right != nil {
		rDif, rMax, rMin := dps(n.Right)
		dif = mins(dif, rDif)
		dif = mins(dif, abs(rMin-n.Val))
		max = rMax
	}
	return dif, max, min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mins(a, b int) int {
	if a < b {
		return a
	}
	return b
}
