package main

// Input: heights = [2,1,5,6,2,3]
// Output: 10
func largestRectangleArea(heights []int) int {
	stack := make([]int, len(heights))
	top := -1
	maxArea := 0
	for i := range heights {
		if top == -1 || heights[i] > heights[stack[top]] {
			top++
			stack[top] = i
		} else {
			for top > -1 {
				if heights[i] < heights[stack[top]] {
					maxArea = max(maxArea, (i-stack[top])*heights[stack[top]])
					top--
					continue
				}
				if heights[i] == heights[stack[top]] {
					top--
					continue
				}
				break
			}
			top++
			heights[stack[top]] = heights[i]
		}
	}
	for top > -1 {
		maxArea = max(maxArea, (len(heights)-stack[top])*heights[stack[top]])
		top--
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
