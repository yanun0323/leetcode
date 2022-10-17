package main

// BUG: [1531] Stuck
func getLengthOfOptimalCompression(s string, k int) int {

	var dfs func(last byte, count int, s []byte, k int) int
	dfs = func(last byte, count int, s []byte, k int) int {
		if len(s) == 0 {
			if count > 0 {
				return 1 + numChar(count)
			}
			return 0
		}
		first := 0

		if s[0] == last {
			first = dfs(last, count+1, s[1:], k)
		} else {
			first = getCount(count+1) + dfs(s[0], 0, s[1:], k)
		}

		if len(s) == 1 {
			return first
		}

		second := 0

		if s[1] == last {
			second = dfs(last, count+1, s[1:], k-1)
		} else {
			second = getCount(count+1) + dfs(s[1], 0, s[2:], k-1)
		}
		return min(first, second)
	}

	return dfs(' ', 0, []byte(s), k)
}

func getCount(count int) int {
	if count == 0 {
		return 0
	}
	return 1 + numChar(count)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numChar(x int) int {
	if x == 1 {
		return 0
	}
	count := 1
	for x >= 10 {
		x /= 10
		count++
	}
	return count
}
