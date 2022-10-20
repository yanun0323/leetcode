package main

import (
	"fmt"
	"sync"
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
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			print(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("complete")
}

func print(i int) {
	fmt.Println("number: ", i)
}
