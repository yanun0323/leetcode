package main

// Input: nums = [4,5,6,7,0,1,2]
// Output: 0
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums)>>1 - (len(nums)+1)%2
	if nums[mid] > nums[len(nums)-1] {
		return findMin(nums[mid+1:])
	}
	return findMin(nums[:mid+1])
}
