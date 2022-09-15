package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return recursive(l1, l2, 0)
}

func recursive(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	node := &ListNode{Val: carry}

	if l1 != nil {
		node.Val += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		node.Val += l2.Val
		l2 = l2.Next
	}

	node.Next = recursive(l1, l2, node.Val/10)
	node.Val %= 10
	return node
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	carry := 0
	current := head
	var node *ListNode
	for l1 != nil || l2 != nil || carry > 0 {
		node = &ListNode{Val: carry}
		current.Next = node
		current = current.Next

		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}
		carry = node.Val / 10
		node.Val %= 10
	}
	return head.Next
}

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node := &ListNode{}
	var nextL1 *ListNode

	if l1 != nil {
		node.Val += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		node.Val += l2.Val
		l2 = l2.Next
	}

	if l1 != nil {
		nextL1 = &ListNode{
			Val:  l1.Val + node.Val/10,
			Next: l1.Next,
		}
	}

	if l1 == nil && node.Val/10 > 0 {
		nextL1 = &ListNode{Val: node.Val / 10}
	}

	node.Next = addTwoNumbers(nextL1, l2)
	node.Val %= 10
	return node
}
