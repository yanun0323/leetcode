package main

type MinStack struct {
	Last *Node
}

func Constructor() MinStack {
	return MinStack{
		Last: nil,
	}
}

func (this *MinStack) Push(val int) {
	if this.Last == nil || val < *this.Last.Min {
		this.Last = NewNode(this.Last, &val, &val)
		return
	}
	this.Last = NewNode(this.Last, &val, this.Last.Min)
}

func (this *MinStack) Pop() {
	if this.Last != nil {
		this.Last = this.Last.Prev
	}
}

func (this *MinStack) Top() int {
	if this.Last == nil {
		return 0
	}
	return *this.Last.V
}

func (this *MinStack) GetMin() int {
	return *this.Last.Min
}

type Node struct {
	Prev *Node
	V    *int
	Min  *int
}

func NewNode(prev *Node, v *int, min *int) *Node {
	return &Node{
		Prev: prev,
		V:    v,
		Min:  min,
	}
}
