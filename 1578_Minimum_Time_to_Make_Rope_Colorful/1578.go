package main

// BUG: [1578] Passed but timeout!
func minCost(colors string, neededTime []int) int {
	result := 0
	max := neededTime[0]
	for i := 1; i < len(colors); i++ {
		if colors[i] != colors[i-1] {
			max = neededTime[i]
			continue
		}

		if neededTime[i] > max {
			result += max
			max = neededTime[i]
			continue
		}

		result += neededTime[i]
	}
	return result
}
