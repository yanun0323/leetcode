package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	INF := math.MaxInt32
	changeCount := make([]int, amount+1)
	for i := range changeCount {
		changeCount[i] = INF
	}
	changeCount[0] = 0
	for money := 0; money <= amount; money++ {
		for _, coin := range coins {
			if money < coin {
				continue
			}
			changeCount[money] = min(changeCount[money-coin]+1, changeCount[money])
		}
	}
	if changeCount[amount] == INF {
		return -1
	}
	return changeCount[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
