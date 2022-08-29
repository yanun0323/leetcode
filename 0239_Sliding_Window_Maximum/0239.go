package main

// TODO: [Yanun] Refactor
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	answer := make([]int, 0, len(nums)-(k-1))
	max := nums[0]
	lp := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		if i == k-1 {
			answer = append(answer, max)
			continue
		}
		if i < k {
			continue
		}
		if nums[lp] == max {
			p := lp + 1
			max = nums[p]
			for ; p < i+1; p++ {
				if nums[p] > max {
					max = nums[p]
				}
			}
		}
		answer = append(answer, max)
		lp++
	}
	return answer
}
