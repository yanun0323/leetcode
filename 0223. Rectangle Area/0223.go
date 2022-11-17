package main

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	return abs(ax2-ax1)*abs(ay2-ay1) + abs(bx2-bx1)*abs(by2-by1) -
		getRepeated(ax1, ax2, bx1, bx2)*getRepeated(ay1, ay2, by1, by2)
}

func getRepeated(a1, a2, b1, b2 int) int {
	if a1 > b1 {
		return getRepeated(b1, b2, a1, a2)
	}

	if a2 <= b1 {
		return 0
	}

	if a2 >= b2 {
		return abs(b2 - b1)
	}

	return abs(a2 - b1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
