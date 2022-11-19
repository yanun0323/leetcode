package main

func countSubstrings(s string) int {
	count := len(s)
	result := 0
	countPalindromic := func(l, r int) {
		for l >= 0 && r < count {
			if s[l] != s[r] {
				break
			}
			result++
			l--
			r++
		}
	}
	for i := range s {
		countPalindromic(i, i)
		if i > 0 {
			countPalindromic(i-1, i)
		}
	}
	return result
}
