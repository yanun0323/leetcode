package main

func permute(nums []int) [][]int {
	result := [][]int{}
	used := map[int]bool{}
	found := make([]int, len(nums))
	var bt func(index int)
	bt = func(index int) {
		if index == len(found) {
			copied := make([]int, len(found))
			copy(copied, found)
			result = append(result, copied)
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			found[index] = nums[i]
			bt(index + 1)
			delete(used, i)
		}
	}
	bt(0)
	return result
}
