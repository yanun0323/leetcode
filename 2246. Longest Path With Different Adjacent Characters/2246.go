package main

type Node struct {
	Char  byte
	Left  *Node
	Right *Node
}

type Candidate struct {
	Chars map[byte]bool
}

func NewCandidate(n *Node) *Candidate {
	return &Candidate{
		Chars: map[byte]bool{n.Char: true},
	}
}

func longestPath(parent []int, s string) int {
	root := parseTree(parent, s)
	q := []*Candidate{}

	leafs := []*Node{}
	var pushLeafs func(n *Node)
	pushLeafs = func(n *Node) {
		if n == nil {
			leafs = append(leafs, n)
			return
		}
		pushLeafs(n.Left)
		pushLeafs(n.Right)
	}
	pushLeafs(root)
	for _, n := range leafs {
		q = append(q, NewCandidate(n))
	}

	return 0
}

func parseTree(parents []int, s string) *Node {
	nodes := make([]*Node, len(parents))
	for i, p := range parents {
		node := &Node{
			Char: s[i],
		}
		nodes[i] = node
		parent := nodes[p]
		if parent == nil {
			continue
		}
		if parent.Left == nil {
			parent.Left = node
			continue
		}
		parent.Right = node
	}
	return nodes[0]
}
