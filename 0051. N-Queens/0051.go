package main

const (
	_Q = 'Q'
	_X = '.'
)

func solveNQueens(n int) [][]string {
	result := [][]string{}

	var m = map[int]string{}
	for i := 0; i < n; i++ {
		m[i] = CreateStr(i, n)
	}

	current := []string{}
	used := make([]map[int]int, n)
	for k := range used {
		used[k] = map[int]int{}
	}
	var bt func(int)
	bt = func(layer int) {
		if layer == n {
			c := make([]string, n)
			copy(c, current)
			result = append(result, c)
			return
		}

		for i := 0; i < n; i++ {
			if used[layer][i] > 0 {
				continue
			}

			newUsed := CreateUsed(layer, i, n)
			for idx := range newUsed {
				for k := range newUsed[idx] {
					used[layer+idx][k]++
				}
			}
			current = append(current, m[i])
			bt(layer + 1)
			current = current[:len(current)-1]

			for idx := range newUsed {
				for k := range newUsed[idx] {
					used[layer+idx][k]--
				}
			}
		}
	}
	bt(0)
	return result
}

func CreateStr(target, n int) string {
	b := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		switch i {
		case target:
			b = append(b, _Q)
		default:
			b = append(b, _X)
		}
	}
	return string(b)
}

func CreateUsed(layer, i, n int) []map[int]bool {
	result := []map[int]bool{}
	l, r := i, i
	for layer < n {
		ins := map[int]bool{l: true, r: true, i: true}
		result = append(result, ins)
		l--
		r++
		layer++
	}
	return result
}
