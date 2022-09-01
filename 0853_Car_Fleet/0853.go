package main

// target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
// TODO: [Yanun] 
func carFleet(target int, position []int, speed []int) int {
	stack := make([]int, 0, len(position))
	temp := make([]int, 0, len(position))
	for i := range position {
		speed[i] = (target - position[i]) / speed[i]
	}
	for i := range position {
		if len(stack) == 0 {
			stack = append(stack, position[i])
			continue
		}

		if stack[len(stack)-1]
	}
}
