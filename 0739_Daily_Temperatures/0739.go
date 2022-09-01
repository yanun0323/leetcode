package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	top := -1
	for i := len(temperatures) - 1; i >= 0; i-- {
		for top > -1 && temperatures[i] >= temperatures[stack[top]] {
			top--
		}
		if top > -1 {
			result[i] = stack[top] - i
		}
		top++
		stack[top] = i
	}
	return result
}

func dailyTemperatures2(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	top := -1
	for i := 0; i < len(temperatures); i++ {
		for top > -1 && temperatures[i] > temperatures[stack[top]] {
			result[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}
	return result
}
