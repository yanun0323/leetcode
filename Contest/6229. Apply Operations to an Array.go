package main

func applyOperations(nums []int) []int {
	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[i-1] == nums[i] {
			nums[i-1] *= 2
			nums[i] = 0
		}
	}

	for s, f := 0, 0; f < l; s++ {
		if nums[s] != 0 {
			continue
		}
		f = s
		for nums[f] == 0 {
			f++
			if f == l {
				return nums
			}
		}
		nums[s] = nums[f]
		nums[f] = 0
	}

	return nums
}
