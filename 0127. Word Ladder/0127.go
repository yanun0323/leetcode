package main

// FIXME: [0127] Timeout Passed
func ladderLength(beginWord string, endWord string, wordList []string) int {
	neighbor := map[string][]string{}
	for i, s1 := range wordList {
		for _, s2 := range wordList[i+1:] {
			if connectable(s1, s2) {
				neighbor[s1] = append(neighbor[s1], s2)
				neighbor[s2] = append(neighbor[s2], s1)
			}
		}
	}

	q := []string{}
	been := map[string]bool{}
	for _, s := range wordList {
		if connectable(beginWord, s) {
			q = append(q, s)
			been[s] = true
		}
	}

	layer := 1
	for len(q) > 0 {
		layer++
		for count := len(q); count > 0; count-- {
			if q[0] == endWord {
				return layer
			}
			for _, s := range neighbor[q[0]] {
				if been[s] {
					continue
				}
				been[s] = true
				q = append(q, s)
			}
			q = q[1:]
		}
	}
	return 0
}

func connectable(a, b string) bool {
	repeated := false
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if repeated {
			return false
		}
		repeated = true
	}
	return repeated
}
