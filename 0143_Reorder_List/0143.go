package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,3,4,5]
// Output: [1,5,2,4,3]

// Input: head = [1,2,3,4]

// stack  321
// tail 45
// head 15 24 3
func reorderList2(head *ListNode) {
	stack := make([]*ListNode, 0, 10)

	for take := head; take != nil; take = take.Next {
		stack = append(stack, take)
	}

	var tail *ListNode
	for {
		tail = head.Next
		if head == stack[len(stack)-1] {
			head.Next = nil
			break
		}
		head.Next = stack[len(stack)-1]
		if stack[len(stack)-1] == tail {
			stack[len(stack)-1].Next = nil
			break
		}
		if stack[len(stack)-1] == tail {
			stack[len(stack)-1].Next = nil
			break
		}
		stack[len(stack)-1].Next = tail
		stack = stack[:len(stack)-1]
		head = head.Next.Next
	}
}

func reorderList(head *ListNode) {
	slow := &ListNode{}
	fast := slow
	slow.Next = head

	// find middle
	// 1234
	// 123
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// tail tail
	tail := slow.Next
	slow.Next = nil
	reverted := revertList(nil, tail)

	// merge two list
	result := &ListNode{}
	for head != nil || reverted != nil {
		result.Next = head
		head = head.Next

		result = result.Next

		if reverted != nil {
			result.Next = reverted
			reverted = reverted.Next
			result = result.Next
		}

	}
	head = result.Next
}

func revertList(last, head *ListNode) *ListNode {
	if head == nil {
		return last
	}
	tail := head.Next
	head.Next = last
	return revertList(head, tail)
}
