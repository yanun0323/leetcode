package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := 1 << 31
	for i := range nums {
		gap := target - nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			diff := gap - nums[l] - nums[r]
			if diff == 0 {
				return target
			}

			if abs(minDiff) > abs(diff) || (abs(minDiff) == abs(diff) && diff > 0) {
				minDiff = diff
			}

			if diff > 0 {
				l++
			} else {
				r--
			}
		}
	}
	return target - minDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
