package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	m := make([]int, n)
	for _, edge := range edges {
		m[edge[1]]++
	}

	min := n
	set := []int{}
	for i, v := range m {
		if v > min {
			continue
		}

		if v == min {
			set = append(set, i)
			continue
		}

		min = v
		set = []int{i}
	}
	return set
}
