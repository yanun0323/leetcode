package main

// FIXME: [0907] 22.11.25 Not the expected method
const mod int = 1e9 + 7

func sumSubarrayMins(arr []int) int {
	result := 0
	for n := len(arr) - 1; n >= 0; n-- {
		result += arr[n]
		for i := 0; i < n; i++ {
			result += arr[i]
			if arr[i+1] < arr[i] {
				arr[i] = arr[i+1]
			}
		}
	}
	return result % mod
}
