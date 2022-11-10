package main

func removeDuplicates(s string) string {
	b := []byte(s)

	for i := 0; i < len(b)-1; {
		if b[i] == b[i+1] {
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
