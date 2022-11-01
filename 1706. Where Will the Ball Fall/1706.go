package main

func findBall(grid [][]int) []int {
	current := make([]int, 0, len(grid[0]))
	for i := range grid[0] {
		current = append(current, i)
	}

	for layer := range grid {
		for i := range current {
			if current[i] == -1 {
				continue
			}
			direction := grid[layer][current[i]]
			neighbor := current[i] + direction
			if neighbor == len(grid[0]) || neighbor < 0 || grid[layer][neighbor] != direction {
				current[i] = -1
				continue
			}
			current[i] += direction
		}
	}
	return current
}
