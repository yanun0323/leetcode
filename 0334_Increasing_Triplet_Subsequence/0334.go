package main

import "math"

// FIXME: [0334] MD
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	smallest, medium := nums[0], math.MaxInt
	for i := range nums {
		switch n := nums[i]; {
		case n < smallest:
			smallest = n
		case n > medium:
			return true
		case n > smallest && n < medium:
			medium = n
		}
	}
	return false
}
