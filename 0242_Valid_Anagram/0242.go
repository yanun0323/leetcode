package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
		count[int(t[i])-int('a')]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}
