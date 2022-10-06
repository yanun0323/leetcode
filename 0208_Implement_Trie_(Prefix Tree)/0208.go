package main

type TrieNode struct {
	stored map[byte]*TrieNode
	isEnd  bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		stored: map[byte]*TrieNode{},
		isEnd:  false,
	}
}

func (n *TrieNode) Contain(b byte) bool {
	return n.stored[b] != nil
}

func (n *TrieNode) Set(b byte) {
	n.stored[b] = NewTrieNode()
}

func (n *TrieNode) Get(b byte) *TrieNode {
	return n.stored[b]
}

func (n *TrieNode) SetEnd() {
	n.isEnd = true
}

func (n *TrieNode) IsEnd() bool {
	return n.isEnd
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for i := range word {
		if !current.Contain(word[i]) {
			current.Set(word[i])
		}
		current = current.Get(word[i])
	}
	current.SetEnd()
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for i := range word {
		if !current.Contain(word[i]) {
			return false
		}
		current = current.Get(word[i])
	}
	return current.IsEnd()
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for i := range prefix {
		if !current.Contain(prefix[i]) {
			return false
		}
		current = current.Get(prefix[i])
	}
	return true
}
