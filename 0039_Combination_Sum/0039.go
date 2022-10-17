package main

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	ans := []int{}

	var bt func(left, index int)
	bt = func(left, index int) {
		if left < 0 {
			return
		}

		if left == 0 {
			copied := make([]int, len(ans))
			copy(copied, ans)
			result = append(result, copied)
			return
		}

		for i := index; i < len(candidates); i++ {
			ans = append(ans, candidates[i])
			bt(left-candidates[i], i)
			ans = ans[:len(ans)-1]
		}
	}

	bt(target, 0)
	return result
}
