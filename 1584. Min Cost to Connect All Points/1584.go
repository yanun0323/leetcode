package main

import "container/heap"

// BUG: [1584] Stuck Passed
type Edge struct {
	a, b     int
	distance int
}

func minCostConnectPoints(p [][]int) int {
	h := &Heap{}
	heap.Init(h)
	for i := range p {
		for j := range p[i+1:] {
			heap.Push(h, Edge{a: i, b: j, distance: getDistance(p[i], p[j])})
		}
	}

	parent := make([]int, h.Len())
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			return find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) bool {
		aP := find(a)
		bP := find(b)
		if aP == bP {
			return false
		}
		parent[aP] = bP
		return true
	}

	result := 0
	for h.Len() > 0 {
		elem := heap.Pop(h).(Edge)
		if union(elem.a, elem.b) {
			result += elem.distance
		}
	}
	return result
}

func getDistance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Heap struct {
	edges []Edge
}

func (h *Heap) Len() int {
	return len(h.edges)
}
func (h *Heap) Less(i, j int) bool {
	return h.edges[i].distance < h.edges[j].distance
}
func (h *Heap) Swap(i, j int) {
	h.edges[i], h.edges[j] = h.edges[j], h.edges[i]
}
func (h *Heap) Push(x interface{}) {
	h.edges = append(h.edges, x.(Edge))
}
func (h *Heap) Pop() interface{} {
	x := h.edges[h.Len()-1]
	h.edges = h.edges[:h.Len()-1]
	return x
}
