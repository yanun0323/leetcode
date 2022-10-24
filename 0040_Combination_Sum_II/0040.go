package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	found := []int{}
	var bt func(int)
	bt = func(first int) {
		if target < 0 {
			return
		}

		if target == 0 {
			co := make([]int, len(found))
			copy(co, found)
			result = append(result, co)
			return
		}

		for i := first; i < len(candidates); i++ {
			if i != first && candidates[i] == candidates[i-1] {
				continue
			}
			num := candidates[i]
			found = append(found, num)
			target -= num
			bt(i + 1)
			target += num
			found = found[:len(found)-1]
		}
	}
	bt(0)
	return result
}
