package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow := dummy
	fast := dummy
	var next *ListNode
	step := 0
	for {
		for step < k && fast.Next != nil {
			fast = fast.Next
			step++
		}
		if step < k {
			return dummy.Next
		}
		step = 0
		next = slow.Next
		slow.Next = reverseList(k-1, fast.Next, slow.Next)
		slow = next
		fast = next
	}
}

func reverseList(time int, prev, head *ListNode) *ListNode {
	if time == 0 {
		head.Next = prev
		return head
	}
	next := head.Next
	head.Next = prev
	return reverseList(time-1, head, next)
}
