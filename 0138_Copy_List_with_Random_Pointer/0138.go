package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldToNew := make(map[*Node]*Node, 0)
	copiedHead := copyNext(head, oldToNew)
	current := copiedHead
	for current != nil {
		if current.Random == nil {
			current = current.Next
			continue
		}
		current.Random = oldToNew[current.Random]
		current = current.Next
	}
	return copiedHead
}

func copyNext(head *Node, pair map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{
		Val:    head.Val,
		Next:   copyNext(head.Next, pair),
		Random: head.Random,
	}
	pair[head] = newHead
	return newHead
}
