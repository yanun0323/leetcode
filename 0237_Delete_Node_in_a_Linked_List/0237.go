package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func deleteNode2(node *ListNode) {
	var shiftNode func(prev, node *ListNode)
	shiftNode = func(prev, node *ListNode) {
		if node.Next == nil {
			prev.Next = nil
			return
		}

		node.Val = node.Next.Val
		shiftNode(node, node.Next)
	}
	shiftNode(nil, node)
}
