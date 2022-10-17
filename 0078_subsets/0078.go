package main

// FIXME: [0078] Timeout Passed
func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	ans := make([]int, 0, len(nums))
	var bt func(sli []int)
	bt = func(sli []int) {
		if len(sli) == 0 {
			return
		}

		for i := range sli {
			ans = append(ans, sli[i])
			l := len(ans)
			copied := make([]int, len(ans))
			copy(copied, ans)
			result = append(result, copied)
			bt(sli[i+1:])
			ans = ans[:l-1]
		}
	}

	bt(nums)
	return result
}

func subsets2(nums []int) [][]int {
	var dfs func(count int, sli []int) [][]int
	dfs = func(count int, sli []int) [][]int {
		if count == 0 {
			return nil
		}

		result := [][]int{}
		for i := 0; i <= len(sli)-count; i++ {
			if count == 1 {
				result = append(result, append([]int{sli[i]}))
				continue
			}
			for _, v := range dfs(count-1, sli[i+1:]) {
				result = append(result, append([]int{sli[i]}, v...))
			}
		}

		return result
	}

	ans := [][]int{}
	for count := len(nums); count > 0; count-- {
		ans = append(ans, dfs(count, nums)...)
	}
	return append(ans, []int{})
}
