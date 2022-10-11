package main

import "container/heap"

type Heap struct {
	m []int
}

func (h *Heap) Len() int {
	return len(h.m)
}
func (h *Heap) Less(i, j int) bool {
	return h.m[i] > h.m[j]
}
func (h *Heap) Swap(i, j int) {
	h.m[i], h.m[j] = h.m[j], h.m[i]
}
func (h *Heap) Pop() interface{} {
	p := h.m[h.Len()-1]
	h.m = h.m[:h.Len()-1]
	return p
}
func (h *Heap) Push(x interface{}) {
	h.m = append(h.m, x.(int))
}

func NewHeap(list []int) *Heap {
	return &Heap{
		m: list,
	}
}

func lastStoneWeight(stones []int) int {
	h := NewHeap(stones)
	heap.Init(h)

	for h.Len() > 1 {
		big := heap.Pop(h).(int)
		small := heap.Pop(h).(int)
		if big != small {
			heap.Push(h, big-small)
			heap.Fix(h, h.Len()-1)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
