package main

func solve(board [][]byte) {
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] != 'O' {
			return
		}
		board[r][c] = '*'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := range board {
		dfs(i, 0)
		dfs(i, len(board[0])-1)
	}

	for i := range board[0] {
		dfs(0, i)
		dfs(len(board)-1, i)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}
