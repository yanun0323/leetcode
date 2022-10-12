package main

func removeDuplicates(nums []int) int {
	count := len(nums)
	if count < 2 {
		return count
	}
	for i := 1; i < count; i++ {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i-1], nums[i:]...)
			count--
			i--
		}
	}
	return len(nums)
}
