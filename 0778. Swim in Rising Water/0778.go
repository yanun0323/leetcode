package main

import "container/heap"

// BUG: [0778] 22.11.11 Stuck

func swimInWater(g [][]int) int {
	h := &Heap{}
	heap.Init(h)
	outOfBounds := func(x, y int) bool {
		return x < 0 || y < 0 || x >= len(g) || y >= len(g[0])
	}
	isNeighbor := func(a, b Node) bool {
		dx, dy := a.X-b.X, a.Y-b.Y
		return (dx == 1 && dy == 0) || (dx == -1 && dy == 0) || (dx == 0 && dy == 1) || (dx == 0 && dy == -1)
	}
	tryPush := func(x, y int) {
		if outOfBounds(x, y) || g[x][y] < 0 {
			return
		}
		heap.Push(h, Node{X: x, Y: y, Val: g[x][y]})
		g[x][y] = -1
	}

	result := g[0][0]
	prev := Node{X: 0, Y: 0, Val: g[0][0]}
	g[0][0] = -1
	heap.Push(h, prev)
	for h.Len() > 0 {
		cur := heap.Pop(h).(Node)
		if isNeighbor(prev, cur) {
			result = max(result, cur.Val)
		} else {
			result = min(result, cur.Val)
		}
		if cur.X == len(g)-1 && cur.Y == len(g[0])-1 {
			return max(result, cur.Val)
		}
		tryPush(cur.X-1, cur.Y)
		tryPush(cur.X, cur.Y-1)
		tryPush(cur.X+1, cur.Y)
		tryPush(cur.X, cur.Y+1)
		prev = cur
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Node struct {
	X, Y, Val int
}

type Heap struct {
	data []Node
}

func (h *Heap) Len() int {
	return len(h.data)
}
func (h *Heap) Less(i, j int) bool {
	return h.data[i].Val < h.data[j].Val
}
func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *Heap) Push(x interface{}) {
	h.data = append(h.data, x.(Node))
}
func (h *Heap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return x
}
