package main

var isVowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

func reverseVowels(s string) string {
	b := []byte(s)
	l, r := 0, len(b)-1
	for l < r {
		if !isVowels[b[l]] {
			l++
		}
		if !isVowels[b[r]] {
			r--
		}
		if isVowels[b[l]] && isVowels[b[r]] {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}
	return string(b)
}
