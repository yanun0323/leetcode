package main

import "container/heap"

func topKFrequent(words []string, k int) []string {
	hash := map[string]int{}
	for i := range words {
		hash[words[i]]++
	}

	h := &Heap{}
	heap.Init(h)
	for k, v := range hash {
		heap.Push(h, Word{
			word:      k,
			frequency: v,
		})
	}

	result := make([]string, 0, h.Len())
	for k > 0 && h.Len() > 0 {
		result = append(result, heap.Pop(h).(Word).word)
		k--
	}
	return result
}

type Word struct {
	word      string
	frequency int
}

type Heap struct {
	words []Word
}

func (h *Heap) Len() int {
	return len(h.words)
}
func (h *Heap) Less(i, j int) bool {
	if h.words[i].frequency == h.words[j].frequency {
		return dictionarySort(h.words[i].word, h.words[j].word)
	}
	return h.words[i].frequency > h.words[j].frequency
}
func (h *Heap) Swap(i, j int) {
	h.words[i], h.words[j] = h.words[j], h.words[i]
}
func (h *Heap) Push(x interface{}) {
	h.words = append(h.words, x.(Word))
}
func (h *Heap) Pop() interface{} {
	x := h.words[h.Len()-1]
	h.words = h.words[:h.Len()-1]
	return x
}
func dictionarySort(a, b string) bool {
	count := min(len(a), len(b))
	for c := 0; c < count; c++ {
		if a[c] == b[c] {
			continue
		}
		return a[c] < b[c]
	}
	return len(a) < len(b)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
