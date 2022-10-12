package main

import "sort"

// BUG: [0976]
// Use sort...

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		if isTriangle(nums[i], nums[i+1], nums[i+2]) {
			ans = max(ans, nums[i]+nums[i+1]+nums[i+2])
		}
	}
	return ans
}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
