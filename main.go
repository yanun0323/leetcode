package main

import (
	"container/heap"
	"fmt"
)

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
func (h *Heap) Pop() interface{} {
	p := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return p
}
func (h *Heap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func main() {
	h := &Heap{data: []int{}}
	heap.Init(h)

	heap.Push(h, 5)
	heap.Push(h, 6)
	heap.Push(h, 7)
	heap.Push(h, 1)
	heap.Push(h, 10)

	fmt.Println(h.data)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
