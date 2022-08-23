package main

func maxArea(height []int) int {
	lp := 0
	rp := len(height) - 1
	max := 0
	current := 0
	for {
		if lp >= rp {
			return max
		}
		current = (rp - lp) * min(height[lp], height[rp])
		if current > max {
			max = current
		}
		if height[lp] > height[rp] {
			rp--
			continue
		}
		lp++
	}
}

func min(l, r int) int {
	if l < r {
		return l
	}
	return r
}
