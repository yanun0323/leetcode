package main

func numIslands(grid [][]byte) int {
	isLandCount := 0
	var explore func(int, int)
	explore = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		explore(x+1, y)
		explore(x-1, y)
		explore(x, y+1)
		explore(x, y-1)
	}

	for i := range grid {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			isLandCount++
			explore(i, j)
		}
	}

	return isLandCount
}
