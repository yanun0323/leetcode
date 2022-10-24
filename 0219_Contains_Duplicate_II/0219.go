package main

// FIXME: [0219] Timeout Passed
func containsNearbyDuplicate(nums []int, k int) bool {
	position := map[int]int{}
	for i := range nums {
		if index, ok := position[nums[i]]; ok && i-index <= k {
			return true
		}
		position[nums[i]] = i
	}
	return false
}
