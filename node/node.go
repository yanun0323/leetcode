package node

type ListNode struct {
	Val  int
	Next *ListNode
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

func (l *ListNode) Arr() []int {
	result := make([]int, 0, 10)
	current := l
	for {
		if current == nil {
			break
		}
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
