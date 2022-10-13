package main

// BUG: [0621]
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	counts := map[byte]int{} // To keep track of cardinality of tasks
	maxCount := 0            // Count the highest occurring task

	for _, task := range tasks { // Handles populating counts and maxCount
		counts[task] += 1
		if counts[task] > maxCount {
			maxCount = counts[task]
		}
	}

	res := ((maxCount - 1) * n) + maxCount - 1 // ((maxCount - 1) * n) Represents the open slots after each task that occurs most often... Add maxCount after to handle the actual slots in which most occurring task runs
	for _, count := range counts {             // Need to add an additional slot for each task that occurs the same number of times as the max occurring one - because that's how many trailing slots will be taken
		if count == maxCount {
			res += 1
		}
	}

	if res < len(tasks) { // If the final result is less than the number of tasks (because n is small compared to maxCount) then just take the length of the tasks
		res = len(tasks)
	}
	return res
}
