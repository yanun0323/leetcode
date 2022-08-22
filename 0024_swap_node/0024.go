package main

import (
	. "main/node"
)

/*
Definition for singly-linked list.

 type ListNode struct {
  	 Val int
  	 Next *ListNode
 }
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	right := head.Next
	head.Next = swapPairs(head.Next.Next)
	right.Next = head

	return right
}
