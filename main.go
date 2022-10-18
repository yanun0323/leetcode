package main

import (
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
	fmt.Println(int('0'))
}
