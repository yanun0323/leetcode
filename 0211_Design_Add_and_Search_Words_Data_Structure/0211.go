package main

// BUG: [0211] Stuck!
type WordNode struct {
	stored map[byte]*WordNode
	isEnd  bool
}

func (n *WordNode) Contain(b byte) bool {
	return n.stored[b] != nil
}

func (n *WordNode) GetNode(b byte) *WordNode {
	return n.stored[b]
}

func NewWordNode() *WordNode {
	return &WordNode{
		stored: map[byte]*WordNode{},
		isEnd:  false,
	}
}

type WordDictionary struct {
	root *WordNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewWordNode(),
	}
}

func (w *WordDictionary) AddWord(word string) {
	current := w.root
	for i := range word {
		if !current.Contain(word[i]) {
			current.stored[word[i]] = NewWordNode()
		}
		current = current.GetNode(word[i])
	}
	current.isEnd = true
}

func (w *WordDictionary) Search(word string) bool {

	var dfs func(node *WordNode, i int) bool
	dfs = func(node *WordNode, i int) bool {
		if node == nil {
			return false
		}

		if i == len(word) {
			return node.isEnd
		}

		if word[i] != '.' {
			if node.Contain(word[i]) {
				return dfs(node.GetNode(word[i]), i+1)
			}
			return false
		}

		for k := range node.stored {
			if dfs(node.stored[k], i+1) {
				return true
			}
		}
		return false
	}
	return dfs(w.root, 0)
}
