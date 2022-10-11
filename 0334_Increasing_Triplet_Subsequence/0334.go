package main

// [1,5,0,4,1,3]
func increasingTriplet(nums []int) bool {
	stack := [2]int{}
	cache := [2]int{}
	top := -1
	cTop := -1

	insert := func(stack [2]int, top *int, num int) bool {
		switch *top {
		case -1:
			*top++
			stack[*top] = num
		case 0:
		case 1:
		case 2:
		}
		return false
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if top == -1 {
			top++
			stack[top] = nums[i]
			continue
		}

		if top == 0 {
			if nums[i] < stack[top] {
				stack[top] = nums[i]
			} else {
				cache
			}
		}

		if top == 1 {

		}

		for top >= 0 && nums[i] >= stack[top] {
			top--
		}

		if top == 1 {
			return true
		}

	}
	return false
}
