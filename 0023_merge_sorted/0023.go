package main

import (
	. "main/node"
)

func mergeKLists(lists []*ListNode) *ListNode {
	// TODO: 0023
	tree := BinaryTree{}
	for _, list := range lists {
		var current *ListNode = list
		for {
			if current == nil {
				break
			}
			tree.Insert(current.Val)
			current = current.Next
		}
	}
	return NewListNode(tree.Access())
}

func NewListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		current = current.Next
	}
	return head
}

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (n *BinaryTree) Access() []int {
	if n.root == nil {
		return nil
	}
	return n.root.access()
}

func (n *BinaryNode) access() []int {
	result := make([]int, 0, 10)

	if n.left != nil {
		if l := n.left.access(); l != nil {
			result = append(result, l...)
		}
	}

	result = append(result, n.data)

	if n.right != nil {
		if r := n.right.access(); r != nil {
			result = append(result, r...)
		}
	}

	return result
}

func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}
