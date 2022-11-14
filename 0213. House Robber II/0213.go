package main

// BUG: [0213] 22.11.15 Stuck Passed
func rob(houses []int) int {
	if len(houses) == 1 {
		return houses[0]
	}
	return max(linerRob(houses[1:]), linerRob(houses[:len(houses)-1]))
}

func linerRob(houses []int) int {
	noTake, goTake := 0, 0
	maxProfit := 0

	for _, money := range houses {
		took := goTake + money
		if took > maxProfit {
			maxProfit = took
		}

		/* Prepare to move next */
		goTake = noTake
		noTake = maxProfit
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
