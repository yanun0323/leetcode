package main

func partition(s string) [][]string {
	result := [][]string{}
	found := []string{}
	var bt func(first int)
	bt = func(first int) {
		if first >= len(s) {
			co := make([]string, len(found))
			copy(co, found)
			result = append(result, co)
			return
		}

		for i := first; i < len(s); i++ {
			if !IsPalindrome(s[first : i+1]) {
				continue
			}
			found = append(found, s[first:i+1])
			bt(i + 1)
			found = found[:len(found)-1]
		}

	}
	bt(0)
	return result
}

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
