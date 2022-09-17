package main

import "strings"

type Pair struct {
	a, b int
}

func palindromePairs(words []string) [][]int {
	wordMap := make(map[string]int, 0)
	result := make([][]int, 0, 10)
	for i := range words {
		wordMap[words[i]] = i
	}

	pairMap := make(map[Pair]bool, 0)
	for i := range words {
		revered := reverseWord(words[i])
		count := len(revered) + 1
		for j := 0; j < count; j++ {
			// word[i] = aabcde; revered[:j] = edcb; revered[j:] = aa
			if k, exist := wordMap[revered[:j]]; exist && k != i && isPalindrome(revered[j:]) {
				pairMap[Pair{k, i}] = true
			}
			// word[i] = aabcde; revered[j:] =dcbaa; revered[:j] = e
			if k, exist := wordMap[revered[j:]]; exist && k != i && isPalindrome(revered[:j]) {
				pairMap[Pair{i, k}] = true
			}
		}
	}
	for k := range pairMap {
		result = append(result, []int{k.a, k.b})
	}

	return result
}

func isPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	}
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func reverseWord(str string) string {
	b := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		b.WriteByte(str[i])
	}
	return b.String()
}
