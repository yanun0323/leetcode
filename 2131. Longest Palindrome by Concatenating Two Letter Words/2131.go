package main

func longestPalindrome(words []string) int {
	result := 0
	selfCount := 0
	self := false
	m := map[string]int{}
	for i := range words {
		self = isSelfPalindrome(words[i])
		if p := palindrome(words[i]); m[p] > 0 {
			if self {
				selfCount--
			}
			m[p]--
			result += 4
			continue
		}
		if self {
			selfCount++
		}
		m[words[i]]++
	}
	if selfCount > 0 {
		return result + 2
	}
	return result
}

func isSelfPalindrome(s1 string) bool {
	return s1[0] == s1[1]
}

func palindrome(s1 string) string {
	return string([]byte{s1[1], s1[0]})
}
