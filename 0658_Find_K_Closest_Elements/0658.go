package main

func findClosestElements(arr []int, k int, x int) []int {
	for len(arr) > k {
		if abs(arr[0]-x) <= abs(arr[len(arr)-1]-x) {
			arr = arr[:len(arr)-1]
			continue
		}
		arr = arr[1:]
	}
	return arr
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
