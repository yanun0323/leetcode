package main

import "sort"

// FIXME: [0332] Timeout Passed
func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i][1] < tickets[j][1]
	})
	m := map[string][]string{}
	for i := range tickets {
		if m[tickets[i][0]] == nil {
			m[tickets[i][0]] = []string{}
		}
		m[tickets[i][0]] = append(m[tickets[i][0]], tickets[i][1])
	}

	left := len(tickets)
	result := make([]string, 0, left+1)
	var dfs func(string) bool
	dfs = func(from string) bool {
		if left == 0 {
			return true
		}
		for i := len(m[from]); i > 0; i-- {
			left--
			dist := m[from][0]
			m[from] = m[from][1:]
			result = append(result, dist)
			if dfs(dist) {
				return true
			}
			left++
			m[from] = append(m[from], dist)
			result = result[:len(result)-1]
		}
		return false
	}
	result = append(result, "JFK")
	dfs("JFK")

	return result
}
