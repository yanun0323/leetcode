package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	lp := 0
	rp := 0
	last := len(height) - 1
	water := 0
	for {
		rp++
		if rp != lp && height[rp] >= height[lp] {
			water += calculateWater(height, lp, rp)
			for {
				lp++
				if lp == rp {
					break
				}
				water -= height[lp]
			}
		}
		if rp == last {
			break
		}
	}
	for {
		if rp == lp {
			return water
		}
		rp--
		if rp != last && height[rp] >= height[last] {
			water += calculateWater(height, rp, last)
			for {
				last--
				if last == rp {
					break
				}
				water -= height[last]
			}
		}
	}
}

func calculateWater(height []int, lp, rp int) int {
	return (rp - lp - 1) * min(height[lp], height[rp])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
