package main

func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if exist[nums[i]] {
			return true
		}
		exist[nums[i]] = true
	}
	return false
}
