package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool, 128)
	length, max := 0, 0
	for lp, rp := 0, 0; rp < len(s); rp++ {
		if seen[s[rp]] {
			if length > max {
				max = length
			}
			for ; seen[s[rp]]; lp++ {
				seen[s[lp]] = false
				length--
			}
		}
		seen[s[rp]] = true
		length++
	}
	if length > max {
		max = length
	}
	return max
}
