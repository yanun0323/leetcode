package main

import "math"

func shortestPath(grid [][]int, k int) int {
	result := math.MaxInt
	var dfs func(r, c, step, k int)
	dfs = func(r, c, step, k int) {
		if r == 0 && c == 0 {
			result = min(result, step)
			return
		}

		k -= grid[r][c]
		if grid[r][c] > 1 || k < 0 {
			return
		}
		grid[r][c] += 2
		if r > 0 {
			dfs(r-1, c, step+1, k)
		}
		if c > 0 {
			dfs(r, c-1, step+1, k)
		}
		grid[r][c] -= 2
	}
	dfs(len(grid)-1, len(grid[0])-1, 0, k)
	if result == math.MaxInt {
		return -1
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
