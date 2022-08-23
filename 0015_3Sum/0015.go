package sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	exists := map[[3]int]bool{}
	for i := 0; i < len(nums); i++ {
		sum := twoSum(nums, -nums[i], i)
		for i, _ := range sum {
			if exists[sum[i]] {
				continue
			}
			exists[sum[i]] = true
			result = append(result, sum[i][:])
		}
	}
	return result
}

func twoSum(nums []int, target int, skipIndex int) [][3]int {
	lp := skipIndex + 1
	rp := len(nums) - 1
	var result [][3]int
	for {
		if lp >= rp {
			break
		}

		if nums[lp]+nums[rp] > target {
			rp--
			continue
		}

		if nums[lp]+nums[rp] < target {
			lp++
			continue
		}
		result = append(result, [3]int{nums[skipIndex], nums[lp], nums[rp]})
		lp++
	}
	return result
}
