package main

// FIXME: [0212]
const _Empty byte = ' '

func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()
	for i := range words {
		trie.Insert(words[i])
	}

	outOfBounds := func(l, r int) bool {
		return l < 0 || r < 0 || l >= len(board) || r >= len(board[0])
	}

	var find func(parent *TriNode, l, r int, result *[]string)
	find = func(parent *TriNode, l, r int, result *[]string) {
		if outOfBounds(l, r) {
			return
		}

		letter := board[l][r]
		node := parent.stored[letter]
		if node == nil {
			return
		}

		if len(node.word) > 0 {
			*result = append(*result, node.word)
			node.word = ""
		}
		board[l][r] = _Empty
		find(node, l+1, r, result)
		find(node, l-1, r, result)
		find(node, l, r+1, result)
		find(node, l, r-1, result)
		board[l][r] = letter
	}

	result := make([]string, 0, len(words))
	for l := 0; l < len(board); l++ {
		for r := 0; r < len(board[l]); r++ {
			if trie.root.stored[board[l][r]] != nil {
				find(trie.root, l, r, &result)
			}
		}
	}
	return result
}

type Trie struct {
	root *TriNode
}

func NewTrie() Trie {
	return Trie{
		root: NewTriNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for i := range word {
		if current.stored[word[i]] == nil {
			current.stored[word[i]] = NewTriNode()
		}
		current = current.stored[word[i]]
	}
	current.word = word
}

type TriNode struct {
	stored map[byte]*TriNode
	word   string
}

func NewTriNode() *TriNode {
	return &TriNode{
		stored: map[byte]*TriNode{},
		word:   "",
	}
}
