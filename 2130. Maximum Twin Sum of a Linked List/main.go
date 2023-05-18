package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	fp, sp := head, head.Next

	var prev *ListNode
	for sp.Next != nil {
		fp, fp.Next, prev = fp.Next, prev, fp
		sp = sp.Next.Next
	}

	max := fp.Val + fp.Next.Val
	fp = fp.Next.Next
	for prev != nil && fp != nil {
		sum := prev.Val + fp.Val
		prev, fp = prev.Next, fp.Next
		if sum > max {
			max = sum
		}
	}
	return max
}
