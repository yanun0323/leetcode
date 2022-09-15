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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	end := &ListNode{}
	var last *ListNode
	for head.Next != nil {
		if head.Next == end {
			return true
		}
		last = head
		head = head.Next
		last.Next = end
	}
	return false
}
