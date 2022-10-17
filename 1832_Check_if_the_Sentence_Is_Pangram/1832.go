package main

const _CharA byte = 'a'

func checkIfPangram(sentence string) bool {
	count := 0
	charMap := [26]bool{}
	for i := range sentence {
		char := int(sentence[i] - _CharA)
		if !charMap[char] {
			charMap[char] = true
			count++
			if count == 26 {
				return true
			}
		}
	}
	return false
}
