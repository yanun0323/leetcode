package main

// WA, TC: O(n), SC: O(n)
func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, 1)
			continue
		}
		result = append(result, result[i-1]*nums[i-1])
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * right
		right *= nums[i]
	}
	return result
}
