package main

/*
containsNearbyAlmostDuplicate

	i != j,
	abs(i - j) <= indexDiff.
	abs(nums[i] - nums[j]) <= valueDiff, and
*/
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	for l, r := 0, 0; r < len(nums); r++ {
		if r-l > indexDiff {
			l++
		}
		for i := l; i < r; i++ {
			if abs(nums[i]-nums[r]) <= valueDiff {
				return true
			}
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
