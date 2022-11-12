package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l, mid, r := 0, x/4, x/2
	for l < r {
		pow := mid * mid
		if pow == x {
			return mid
		}
		if pow > x {
			l = mid
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}
	return l
}
