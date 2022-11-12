package main

func makeGood(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)-1; {
		if b[i]+32 == b[i+1] || b[i]-32 == b[i+1] {
			b = append(b[:i], b[i+2:]...)
			if i > 0 {
				i--
			}
			continue
		}
		i++
	}
	return string(b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
