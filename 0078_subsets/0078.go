package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0, 10)
	for _, num := range nums {
		start := len(result)
		result = append(result, []int{num})
		for i := 0; i < start; i++ {
			new := result[i]
			new = append(new, num)
			result = append(result, new)
		}
	}
	answer := make([][]int, 0, len(result)+1)
	answer = append(answer, []int{})
	answer = append(answer, result...)
	return answer
}
