package main

func exist(board [][]byte, word string) bool {
	index := 0
	found := false

	var findWord func(int, int)
	findWord = func(r, c int) {
		if found || r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
			return
		}

		if board[r][c] != word[index] {
			return
		}

		if index == len(word)-1 {
			found = true
			return
		}
		board[r][c] = '*'
		index++
		findWord(r+1, c)
		findWord(r-1, c)
		findWord(r, c+1)
		findWord(r, c-1)
		index--
		board[r][c] = word[index]
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				findWord(i, j)
			}
		}
	}
	return found
}
