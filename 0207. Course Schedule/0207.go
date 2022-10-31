package main

// FIXME: [0207] Timeout Passed
func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make([]map[int]bool, numCourses)
	for i := range prerequisites {
		idx := prerequisites[i][1]
		if nodes[idx] == nil {
			nodes[idx] = map[int]bool{}
		}
		nodes[idx][prerequisites[i][0]] = true
	}

	var valid func(course int) bool
	valid = func(course int) bool {
		if nodes[course] == nil {
			return true
		}

		if nodes[course][-1] {
			return false
		}

		nodes[course][-1] = true
		for k := range nodes[course] {
			if k == -1 {
				continue
			}
			if !valid(k) {
				return false
			}
		}
		nodes[course][-1] = false
		nodes[course] = nil
		return true
	}

	for k := range nodes {
		if !valid(k) {
			return false
		}
	}
	return true
}
