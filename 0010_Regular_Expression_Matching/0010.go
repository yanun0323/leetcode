package main

// BUG: [0010]
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if p[len(p)-1] == '*' {
		return isMatch(s, p[:len(p)-2]) || len(s) > 0 && (p[len(p)-2] == '.' || s[len(s)-1] == p[len(p)-2]) && isMatch(s[:len(s)-1], p)
	}

	if len(s) > 0 && (p[len(p)-1] == '.' || s[len(s)-1] == p[len(p)-1]) {
		return isMatch(s[:len(s)-1], p[:len(p)-1])
	}

	return false
}
