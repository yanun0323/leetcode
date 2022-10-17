package main

import "container/heap"

// FIXME: [0295]
type MedianFinder struct {
	right, reversed *Heap
}

func Constructor() MedianFinder {
	r, l := &Heap{}, &Heap{}

	heap.Init(r)
	heap.Init(l)
	return MedianFinder{
		right:    r,
		reversed: l,
	}
}

func (f *MedianFinder) AddNum(num int) {
	heap.Push(f.reversed, -num)
	heap.Push(f.right, -heap.Pop(f.reversed).(int))
	if f.right.Len() > f.reversed.Len() {
		heap.Push(f.reversed, -heap.Pop(f.right).(int))
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if (f.right.Len()+f.reversed.Len())%2 == 1 {
		return float64(-f.reversed.data[0])
	}
	return float64(f.right.data[0]+(-f.reversed.data[0])) * 0.5
}

type Heap struct {
	data []int
}

func (h *Heap) Len() int {
	return len(h.data)
}
func (h *Heap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}
func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *Heap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}
func (h *Heap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}
