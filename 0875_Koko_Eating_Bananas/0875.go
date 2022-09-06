package main

func minEatingSpeed(piles []int, h int) int {
	return findK(1, findMax(piles), piles, &h)
}

func findK(l, r int, piles []int, h *int) int {
	if l > r {
		return l
	}

	mid := (l + r) >> 1
	if canEatAll(mid, piles, h) {
		return findK(l, mid-1, piles, h)
	}
	return findK(mid+1, r, piles, h)
}

func canEatAll(mid int, piles []int, h *int) bool {
	times := 0
	for i := range piles {
		times += (piles[i] + mid - 1) / mid
	}
	return times <= *h
}

func findMax(piles []int) int {
	max := piles[0]
	for i := range piles {
		if piles[i] > max {
			max = piles[i]
		}
	}
	return max
}
