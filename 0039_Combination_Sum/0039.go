package main

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}

	var bt func(left, index int)
	bt = func(left, index int) {
		if left < 0 {
			return
		}

		if left == 0 {
			found := make([]int, len(current))
			copy(found, current)
			result = append(result, found)
			return
		}

		for i := index; i < len(candidates); i++ {
			current = append(current, candidates[i])
			bt(left-candidates[i], i)
			current = current[:len(current)-1]
		}
	}

	bt(target, 0)
	return result
}
