package main

// FIXME: [0645] Timeout Passed
func findErrorNums(nums []int) []int {
	set := make(map[int]bool, len(nums))
	repeat := 0
	sum := 0
	for i := range nums {
		if set[nums[i]] {
			repeat = nums[i]
		}
		set[nums[i]] = true
		sum += nums[i]
	}

	expectedSum := int(float64(1+len(nums)) * (float64(len(nums)) / 2.0))
	num := expectedSum - sum + repeat
	return []int{repeat, num}
}
