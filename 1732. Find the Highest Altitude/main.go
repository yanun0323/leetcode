package main

func largestAltitude(gain []int) int {
	result := 0
	altitude := 0
	for _, n := range gain {
		altitude += n
		result = max(altitude, result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
