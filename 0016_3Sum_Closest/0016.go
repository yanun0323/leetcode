package main

import "sort"

// BUG: [0016] Stuck!
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	smallestDiff := 1 << 31

	for ; len(nums) > 2; nums = nums[1:] {
		l, r := 1, len(nums)-1
		for l < r {
			diff := target - nums[0] - nums[l] - nums[r]
			if diff == 0 {
				return target
			}

			if abs(diff) < abs(smallestDiff) || (abs(diff) == abs(smallestDiff) && diff > smallestDiff) {
				smallestDiff = diff
			}

			if diff > 0 {
				l++
			} else {
				r--
			}
		}
	}
	return target - smallestDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
