package main

func removeElement(nums []int, val int) int {
	count := len(nums)
	for i := 0; i < count; i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			count--
			i--
		}
	}
	return count
}
