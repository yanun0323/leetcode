package main

func maxSlidingWindow(nums []int, k int) []int {
	answer := make([]int, 0, len(nums)-(k-1))
	win := make([]int, 0, k)
	for i := range nums {
		l := len(win) - 1
		for l >= 0 && nums[i] > nums[win[l]] {
			l--
		}
		if len(win) != 0 {
			if l == -1 || win[0] != i-k {
				win = win[:l+1]
			} else {
				win = win[1 : l+1]
			}
		}
		win = append(win, i)
		if i >= k-1 {
			answer = append(answer, nums[win[0]])
		}
	}
	return answer
}
