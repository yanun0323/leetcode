package main

func breakPalindrome(p string) string {
	if len(p) == 1 {
		return ""
	}

	buf := []byte(p)
	for l, r := 0, len(p)-1; l < r; l, r = l+1, r-1 {
		if buf[l] != 'a' {
			buf[l] = 'a'
			return string(buf)
		}
	}

	buf[len(buf)-1] = 'b'
	return string(buf)
}
