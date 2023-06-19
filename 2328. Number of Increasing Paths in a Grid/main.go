package main

type point [2]int

func (p point) x() int {
	return p[1]
}

func (p point) y() int {
	return p[0]
}

func (p point) add(y, x int) point {
	return point{p[0] + y, p[1] + x}
}

func (p point) outOfBounds(grid [][]int) bool {
	return p.y() < 0 || p.y() >= len(grid) || p.x() < 0 || p.x() >= len(grid[0])
}

const (
	mod      = 1_000_000_007
	maxValue = 100_001
)

func countPaths(grid [][]int) int {
	dp := make(map[point]int, len(grid)*len(grid[0]))
	var dfs func(point, int) int
	dfs = func(p point, previous int) int {
		if p.outOfBounds(grid) || previous <= grid[p.y()][p.x()] {
			return 0
		}

		if dp[p] > 0 {
			return dp[p]
		}

		dp[p] = 1
		v := grid[p.y()][p.x()]
		dp[p] = (dp[p] + dfs(p.add(1, 0), v)) % mod
		dp[p] = (dp[p] + dfs(p.add(0, 1), v)) % mod
		dp[p] = (dp[p] + dfs(p.add(-1, 0), v)) % mod
		dp[p] = (dp[p] + dfs(p.add(0, -1), v)) % mod
		return dp[p]
	}

	result := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			result = (result + dfs(point{y, x}, maxValue)) % mod
		}
	}
	return result
}
