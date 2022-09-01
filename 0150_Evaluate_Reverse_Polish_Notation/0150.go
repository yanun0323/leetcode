package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	top := -1

	for i := range tokens {
		switch tokens[i] {
		case "+":
			stack[top-1] = stack[top-1] + stack[top]
			top--
		case "-":
			stack[top-1] = stack[top-1] - stack[top]
			top--
		case "*":
			stack[top-1] = stack[top-1] * stack[top]
			top--
		case "/":
			stack[top-1] = stack[top-1] / stack[top]
			top--
		default:
			top++
			num, _ := strconv.Atoi(tokens[i])
			stack[top] = num
		}
	}
	return stack[top]
}
