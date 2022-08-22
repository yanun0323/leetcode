package main

func longestConsecutive(nums []int) int {
	dic := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] = true
	}

	result := 0
	length := 1
	for n := range dic {
		if dic[n-1] {
			continue
		}
		for i := n + 1; dic[i]; i++ {
			length++
		}
		result = max(result, length)
		length = 1
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
