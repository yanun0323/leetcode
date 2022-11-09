package main

func maximumSubarraySum(nums []int, k int) int64 {
	been := map[int]int{}
	var result int64 = 0
	var sum int64 = 0
	l, r := 0, 0
	for ; r < k; r++ {
		been[nums[r]]++
		sum += int64(nums[r])
	}
	if len(been) == k {
		result = max(sum, result)
	}

	for r < len(nums) {
		sum -= int64(nums[l])
		been[nums[l]]--
		if been[nums[l]] == 0 {
			delete(been, nums[l])
		}
		l++
		sum += int64(nums[r])
		been[nums[r]]++
		r++
		if len(been) == k {
			result = max(sum, result)
		}
	}
	return result
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
