package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// Cache the variable
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func reverseList_DivAndCon(head *ListNode) *ListNode {
	var rest *ListNode
	return DivAndCon(nil, head, rest)
}

func DivAndCon(prev, head, rest *ListNode) *ListNode {
	if head == nil {
		return prev
	}
	rest = head.Next
	head.Next = prev
	return DivAndCon(head, rest, prev)
}

// 1 2
