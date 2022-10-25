package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	cloned := map[int]*Node{}
	var clone func(*Node) *Node
	clone = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if nn, exist := cloned[n.Val]; exist {
			return nn
		}

		copied := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, 0, len(n.Neighbors)),
		}
		cloned[copied.Val] = copied
		for i := range n.Neighbors {
			copied.Neighbors = append(copied.Neighbors, clone(n.Neighbors[i]))
		}
		return copied
	}

	return clone(node)
}
