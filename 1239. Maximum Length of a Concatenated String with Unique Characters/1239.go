package main

const _Invalid byte = '*'

// TODO: [Yanun] Implement me
func maxLength(arr []string) int {
	setMap := make(map[int]map[byte]bool, len(arr))
	result, current := 0, 0
	var bt func(int)
	bt = func(first int) {
		if first >= len(arr) {
			result = max(result, current)
			return
		}
		if setMap[first] == nil {
			setMap[first] = CreateSet(arr[first])
		}
		if setMap[first][_Invalid] {
			result = max(result, current)
			return
		}
		for i := first + 1; i < len(arr); i++ {
			if setMap[i] == nil {
				setMap[i] = CreateSet(arr[i])
			}
			if setMap[i][_Invalid] {
				continue
			}
			good := true
			for char := range setMap[first] {
				if setMap[i][char] {
					good = false
					break
				}
			}
			if good {
				current += len(arr[i])
				bt(i + 1)
				current -= len(arr[i])
			}
		}
		result = max(result, current)
	}
	bt(0)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CreateSet(s string) map[byte]bool {
	set := map[byte]bool{}
	for i := range s {
		if set[s[i]] {
			return map[byte]bool{_Invalid: true}
		}
		set[s[i]] = true
	}
	return set
}
