package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	m1 := [26]int{}
	m2 := [26]int{}

	for i := range word1 {
		m1[int(word1[i]-'a')]++
		m2[int(word2[i]-'a')]++
	}

	count := map[int]int{}
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] && (m1[i] == 0 || m2[i] == 0) {
			return false
		}
		count[m1[i]]++
		count[m2[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
