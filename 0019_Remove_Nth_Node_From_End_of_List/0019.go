package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//
// 123456
// 12345
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	count := 0
	slow := &ListNode{}
	fast := slow
	slow.Next = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		count++
	}
	count *= 2
	if fast == nil { //奇數
		count -= 1
	}

	if n == count {
		return head.Next
	}

	current := head
	r := head.Next
	for i := 0; i < count-n-1; i++ {
		current = current.Next
		if r != nil {
			r = r.Next
		}
	}
	if r == nil {
		return head
	}
	current.Next = r.Next
	return head
}

// 2

// X X X X X

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := &ListNode{}
	fast := slow
	slow.Next = head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast.Next == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
