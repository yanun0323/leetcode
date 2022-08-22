package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	width := right

	for {
		if left == right {
			break
		}
		if height[left] > height[right] {
			max = maxNum(max, width*height[right])
			right -= 1
		} else {
			max = maxNum(max, width*height[left])
			left += 1
		}
		width -= 1
	}
	return max
}

func maxNum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
