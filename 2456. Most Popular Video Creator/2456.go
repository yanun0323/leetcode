package main

import "sort"

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	hash := map[string]*Tuber{}
	list := []*Tuber{}
	for i := range creators {
		if t, exist := hash[creators[i]]; exist {
			t.views += views[i]
			if t.maxView < views[i] {
				t.maxID = ids[i]
				t.maxView = views[i]
			}
			if t.maxView == views[i] && t.maxID > ids[i] {
				t.maxID = ids[i]
			}
			continue
		}

		t := &Tuber{
			creator: creators[i],
			maxID:   ids[i],
			maxView: views[i],
			views:   views[i],
		}
		hash[t.creator] = t
		list = append(list, t)
	}
	result := make([][]string, 0, len(hash))
	sort.Slice(list, func(i, j int) bool {
		return list[i].views > list[j].views
	})
	maxView := list[0].views
	for i := range list {
		if list[i].views != maxView {
			break
		}
		result = append(result, []string{list[i].creator, list[i].maxID})
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Tuber struct {
	creator string
	maxID   string
	maxView int
	views   int
}
