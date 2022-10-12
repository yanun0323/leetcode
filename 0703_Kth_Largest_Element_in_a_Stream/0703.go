package main

import (
	"container/heap"
	"sort"
)

// BUG: [0703]
type Heap struct {
	l []int
}

func (h *Heap) Len() int {
	return len(h.l)
}
func (h *Heap) Less(i, j int) bool {
	return h.l[i] < h.l[j]
}
func (h *Heap) Swap(i, j int) {
	h.l[i], h.l[j] = h.l[j], h.l[i]
}
func (h *Heap) Pop() interface{} {
	p := h.l[h.Len()-1]
	h.l = h.l[:h.Len()-1]

	return p
}
func (h *Heap) Push(x interface{}) {
	h.l = append(h.l, x.(int))
}

func NewHeap(list []int) *Heap {
	return &Heap{
		l: list,
	}
}

type KthLargest struct {
	heap *Heap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	if k < len(nums) {
		nums = nums[:k]
	}
	h := NewHeap(nums)
	heap.Init(h)
	return KthLargest{
		heap: h,
		k:    k,
	}
}

func (kl *KthLargest) Add(val int) int {
	if kl.heap.Len() < kl.k {
		heap.Push(kl.heap, val)
	} else if val > kl.heap.l[0] {
		kl.heap.l[0] = val
		heap.Fix(kl.heap, 0)
	}

	return kl.heap.l[0]
}
