package main

type TriNode struct {
	stored map[byte]*TriNode
	isWord bool
}

func NewTriNode() *TriNode {
	return &TriNode{
		stored: map[byte]*TriNode{},
		isWord: false,
	}
}

type Trie struct {
	root *TriNode
}

func NewTrie() Trie {
	return Trie{
		root: NewTriNode(),
	}
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	charSheet := map[byte][][]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			char := board[i][j]
			charSheet[char] = append(charSheet[char], []int{i, j})
		}
	}

	for i := range words {
		char := words[i][0]
		if len(charSheet[char]) == 0 {
			continue
		}
		found := false
		iter := NewIterator(board, words[i], &found, &result)
		for _, coord := range charSheet[char] {
			iter.Find(coord[0], coord[1])
			if found {
				break
			}
		}
	}
	return result
}

type Iterator struct {
	board  [][]byte
	word   string
	index  int
	found  *bool
	result *[]string
}

func NewIterator(board [][]byte, word string, found *bool, result *[]string) Iterator {
	return Iterator{
		board:  board,
		word:   word,
		index:  0,
		found:  found,
		result: result,
	}
}

func (i Iterator) Find(l, r int) {
	if *i.found {
		return
	}
	if i.index >= len(i.word) {
		*i.found = true
		*i.result = append(*i.result, i.word)
		return
	}

	if l < 0 || r < 0 || l >= len(i.board) || r >= len(i.board[0]) {
		return
	}

	if i.board[l][r] != i.word[i.index] {
		return
	}
	letter := i.board[l][r]
	i.board[l][r] = '.'
	i.index++

	i.Find(l-1, r)
	i.Find(l, r-1)
	i.Find(l+1, r)
	i.Find(l, r+1)
	i.board[l][r] = letter
}
