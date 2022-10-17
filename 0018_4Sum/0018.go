package main

import "sort"

// FIXME: [1046]
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target, k int) [][]int {
	result := [][]int{}

	if k > 2 {
		for i := 0; i < len(nums)-(k-1); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			sums := kSum(nums[i+1:], target-nums[i], k-1)
			for s := range sums {
				sums[s] = append([]int{nums[i]}, sums[s]...)
			}
			result = append(result, sums...)
		}
		return result
	}

	if k == 2 && len(nums) >= 2 {
		l, r := 0, len(nums)-1
		for l < r {
			num := nums[l] + nums[r]
			if num == target {
				result = append(result, []int{nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				continue
			}

			if num < target {
				l++
			} else {
				r--
			}
		}
	}

	return result
}
