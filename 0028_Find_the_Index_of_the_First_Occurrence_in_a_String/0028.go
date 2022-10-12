package main

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if haystack[i] == needle[0] && len(needle) >= len(haystack)-i {
			mismatch := false
			for j := range needle {
				if haystack[i+j] != needle[j] {
					mismatch = true
					break
				}
			}
			if mismatch {
				continue
			}
			return i
		}
	}
	return -1
}
