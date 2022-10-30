package main

// FIXME: [0994] Timeout Passed
type Position struct {
	X, Y int
}

func orangesRotting(grid [][]int) int {
	timePass := 0
	freshCount := 0
	q := []Position{}
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 1:
				freshCount++
			case 2:
				q = append(q, Position{X: i, Y: j})
			}
		}
	}

	tryPush := func(x, y int) {
		if grid[x][y] != 1 { return }
		freshCount--
		grid[x][y] = 2
		q = append(q, Position{X: x, Y: y})
	}

	for len(q) > 0 {
		for size := len(q); size > 0; size-- {
			timePass++
			x, y := q[0].X, q[0].Y

			if x > 0 { tryPush(x-1, y) }
			if y > 0 { tryPush(x, y-1) }
			if x+1 < len(grid) { tryPush(x+1, y) }
			if y+1 < len(grid[0]) { tryPush(x, y+1) }
			q = q[:0]
		}
	}
	if freshCount > 0 { return -1 }
	return timePass
}

func orangesRotting2(grid [][]int) int {
	minusPass := 0
	var infect func(x, y, minusTime int)
	infect = func(x, y, minusTime int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return
		}

		switch grid[x][y] {
		case 0, 2:
			return
		case 1:
			grid[x][y] = 0
		default:
			if grid[x][y] > minusTime {
				return
			}
		}
		infect(x+1, y, minusTime-1)
		infect(x, y+1, minusTime-1)
		infect(x-1, y, minusTime-1)
		infect(x, y-1, minusTime-1)
		grid[x][y] = minusTime
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				infect(i+1, j, -1)
				infect(i, j+1, -1)
				infect(i-1, j, -1)
				infect(i, j-1, -1)
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
			if grid[i][j] < minusPass {
				minusPass = grid[i][j]
			}
		}
	}
	return -minusPass
}
