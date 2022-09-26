package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoList(mergeTwoList(lists[0], lists[1]), mergeKLists(lists[2:]))
}

func mergeTwoList(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l != nil || r != nil {
		if l == nil {
			current.Next = r
			return dummy.Next
		}
		if r == nil {
			current.Next = l
			return dummy.Next
		}
		if l.Val < r.Val {
			current.Next = l
			current = current.Next
			l = l.Next
			continue
		}
		current.Next = r
		current = current.Next
		r = r.Next
	}
	return dummy.Next
}
