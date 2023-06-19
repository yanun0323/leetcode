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
	p1 := [2]int{3, 5}
	p2 := [2]int{1 + 2, 3 + 2}
	m := map[[2]int]int{}
	m[p1] = 10
	fmt.Println(m[p2])
}

func print(i int) {
	fmt.Println("number: ", i)
}
