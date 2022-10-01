package main

import "sort"

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	answer := [][]int{}
	uniqPositions := make(map[int]struct{}, n*2)
	for _, building := range buildings {
		uniqPositions[building[0]] = struct{}{}
		uniqPositions[building[1]] = struct{}{}
	}

	sortedUniqPositions := make([]int, 0, len(uniqPositions))
	for k, _ := range uniqPositions {
		sortedUniqPositions = append(sortedUniqPositions, k)
	}

	sort.Ints(sortedUniqPositions)

	for _, pos := range sortedUniqPositions {
		maxHeight := 0
		for _, building := range buildings {
			if building[0] <= pos && pos < building[1] {
				maxHeight = max(maxHeight, building[2])
			}
		}

		if len(answer) == 0 || answer[len(answer)-1][1] != maxHeight {
			answer = append(answer, []int{pos, maxHeight})
		}
	}
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
