package main

import "strconv"

func summaryRanges(nums []int) []string {
	count := len(nums)
	if count == 0 {
		return []string{}
	}
	if count == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	result := make([]string, 0, count)
	slow, fast := 0, 1
	for fast <= count {
		if fast != count && nums[fast-1]+1 == nums[fast] {
			fast++
			continue
		}

		if slow == fast-1 {
			result = append(result, strconv.Itoa(nums[slow]))
		} else {
			str := strconv.Itoa(nums[slow]) + "->" + strconv.Itoa(nums[fast-1])
			result = append(result, str)
		}
		slow = fast
		fast++
	}
	return result
}
