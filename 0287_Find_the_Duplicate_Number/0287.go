package main

// 龜兔賽跑演算法
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if nums[slow] == nums[fast] {
			break
		}
	}
	slow = 0

	for nums[slow] != nums[fast] {
		slow = nums[slow]
		fast = nums[fast]
	}

	return nums[slow]
}
