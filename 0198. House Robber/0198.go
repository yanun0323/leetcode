package main

func rob(houses []int) int {
	// goTake, noTake,   *   , _
	//         goTake, noTake, *
	noTake, goTake := 0, 0
	maxProfit := 0

	for _, money := range houses {
		took := goTake + money
		maxProfit = max(took, noTake)

		/* Next */
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
