package main

// BUG: [2458] Stuck
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	m := map[int]int{}
	q := []*TreeNode{root}
	single := 0
	layer := 0

	for ; len(q) > 0; layer++ {
		single = 0
		if len(q) == 1 {
			single = 1
		}
		for c := len(q); c > 0; c-- {
			m[q[0].Val] = layer - single
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}
	}
	result := make([]int, 0, len(queries))
	for i := range queries {
		result = append(result, m[queries[i]])
	}
	return result
}
