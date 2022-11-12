package main

import (
	"math"
)

type Flight struct {
	To    int
	Price int
}

type Stop struct {
	Been bool
	Step int
	Cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	m := make(map[int][]Flight, n)
	for _, f := range flights {
		m[f[0]] = append(m[f[0]], Flight{
			To:    f[1],
			Price: f[2],
		})
	}

	stops := make(map[int]*Stop, n)
	var dfs func(int)
	dfs = func(s int) {
		stopS := stops[s]
		stopS.Been = true
		currentStep, currentCost := stopS.Step, stopS.Cost

		// handle neighbor
		currentStep++
		minGoing := s
		minPrice := math.MaxInt
		for len(m[s]) > 0 {
			going := m[s][0]
			m[s] = m[s][1:]
			thisCost := currentCost + going.Price
			targetStop, exist := stops[going.To]
			if exist && targetStop.Been {
				return
			}

			if !exist {
				stops[going.To] = &Stop{
					Been: false,
					Step: currentStep,
					Cost: thisCost,
				}
			} else if targetStop.Step > k || (currentStep <= k && targetStop.Cost > thisCost) {
				targetStop.Step = currentStep
				targetStop.Cost = thisCost
			}

			if going.Price < minPrice {
				minPrice = going.Price
				minGoing = going.To
			}
		}
		if minGoing == s {
			return
		}
		dfs(minGoing)
	}
	stops[src] = &Stop{
		Been: true,
		Step: -1,
	}
	dfs(src)
	dstStop, exist := stops[dst]
	if !exist || dstStop.Step > k {
		return -1
	}
	return dstStop.Cost
}
