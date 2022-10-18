package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{[]int{}}
	current := make([]int, 0, len(nums))
	var bt func(first int)
	bt = func(first int) {
		if first == len(nums) {
			return
		}

		used := map[int]bool{}
		for i := first; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true

			current = append(current, nums[i])
			found := make([]int, len(current))
			copy(found, current)
			result = append(result, found)

			bt(i + 1)
			current = current[:len(current)-1]
		}
	}
	bt(0)
	return result
}
