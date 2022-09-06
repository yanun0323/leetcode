package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(popped))
	match := 0
	for i := range pushed {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
			match++
		}
	}

	return match == len(pushed)
}
