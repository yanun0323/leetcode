package main

// "AABABBA" 1
func characterReplacement(s string, k int) int {
	frequency := [26]int{}
	answer := 0
	max := 0
	for start, end := 0, 0; end < len(s); end++ {
		frequency[s[end]-'A']++
		if frequency[s[end]-'A'] > max {
			max = frequency[s[end]-'A']
		}
		if end-start+1-max > k {
			frequency[s[start]-'A']--
			start++
		}

		if end-start+1 > answer {
			answer = end - start + 1
		}
	}
	return answer
}
