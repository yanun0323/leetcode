package main

import (
	"container/heap"
	"strings"
)

func frequencySort(s string) string {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}
	h := &Heap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, Char{
			char:      k,
			frequency: v,
		})
	}
	b := strings.Builder{}
	for h.Len() > 0 {
		c := heap.Pop(h).(Char)
		for i := 0; i < c.frequency; i++ {
			b.WriteByte(c.char)
		}
	}
	return b.String()
}

type Char struct {
	char      byte
	frequency int
}

type Heap struct {
	data []Char
}

func (h *Heap) Len() int {
	return len(h.data)
}
func (h *Heap) Less(i, j int) bool {
	return h.data[i].frequency > h.data[j].frequency
}
func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *Heap) Push(x interface{}) {
	h.data = append(h.data, x.(Char))
}
func (h *Heap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}
