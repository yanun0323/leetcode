package main

// BUG: [0210] Stuck Passed
type Node struct {
	Require int
	Prev    map[int]bool
}

func NewNode() Node {
	return Node{Prev: map[int]bool{}}
}

func (n *Node) IsEnd() bool {
	return n.Prev == nil
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([]Node, numCourses)
	q := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = NewNode()
		q[i] = i
	}
	for _, pre := range prerequisites {
		graph[pre[0]].Require++
		q[pre[0]] = -1
		graph[pre[1]].Prev[pre[0]] = true
	}
	result := []int{}
	for ; len(q) > 0; q = q[1:] {
		if q[0] == -1 {
			continue
		}
		for i := range graph[q[0]].Prev {
			graph[i].Require--
			if graph[i].Require == 0 {
				q = append(q, i)
			}
		}
		result = append(result, q[0])

	}
	if len(result) != numCourses {
		return nil
	}
	return result
}
