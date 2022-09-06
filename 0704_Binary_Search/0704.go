package main

func search(nums []int, target int) int {
	return DivAndCoq(0, nums, target)
}

func DivAndCoq(index int, nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return index
		}
		return -1
	}
	if nums[len(nums)/2] > target {
		return DivAndCoq(index, nums[:len(nums)/2], target)
	}
	return DivAndCoq(index+len(nums)/2, nums[len(nums)/2:], target)
}
