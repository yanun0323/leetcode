package main

// FIXME: [0684] Timeout Passed
type Node struct {
	Neighbor []int
}

func findRedundantConnection(edges [][]int) []int {
	graph := map[int]Node{}
	isCircle := func(a, b int) bool {
		been := map[int]bool{}
		for q := []int{a}; len(q) > 0; q = q[1:] {
			if q[0] == b {
				return true
			}
			if been[q[0]] {
				continue
			}
			been[q[0]] = true
			for _, i := range graph[q[0]].Neighbor {
				q = append(q, i)
			}
		}
		return false
	}

	var result []int
	for i := range edges {
		pair := edges[i]
		if graph[pair[0]].Neighbor != nil && graph[pair[1]].Neighbor != nil {
			if isCircle(pair[0], pair[1]) {
				result = pair
			}
		}
		n0 := graph[pair[0]]
		n0.Neighbor = append(n0.Neighbor, pair[1])
		graph[pair[0]] = n0
		n1 := graph[pair[1]]
		n1.Neighbor = append(n1.Neighbor, pair[0])
		graph[pair[1]] = n1
	}
	return result
}
