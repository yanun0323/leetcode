package main

// ["c==d","f!=a","b==f","b==c"] d=0  [d!=d]

func equationsPossible(equations []string) bool {
	roots := [26]int{}
	for i := range roots {
		roots[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		for i != roots[i] {
			i = roots[i]
		}
		return i
	}

	var union func(int, int)
	union = func(i, j int) {
		roots[find(i)] = find(j)
	}

	for i := range equations {
		if equations[i][1] == '=' {
			union(int(equations[i][0]-'a'), int(equations[i][3]-'a'))
		}
	}

	for i := range equations {
		if equations[i][1] == '!' && roots[find(int(equations[i][0]-'a'))] == roots[find(int(equations[i][3]-'a'))] {
			return false
		}
	}
	return true
}
