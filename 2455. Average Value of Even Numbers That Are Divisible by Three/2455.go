package main

func averageValue(nums []int) int {
	sum, count := 0, 0
	for i := range nums {
		if nums[i] > 0 && nums[i]%6 == 0 {
			sum += nums[i]
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
