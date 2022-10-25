package main

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	var explore func(x, y int) int
	explore = func(x, y int) int {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
			return 0
		}
		grid[x][y] = 0
		return 1 + explore(x+1, y) + explore(x-1, y) + explore(x, y+1) + explore(x, y-1)
	}

	for i := range grid {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			maxArea = max(maxArea, explore(i, j))
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
