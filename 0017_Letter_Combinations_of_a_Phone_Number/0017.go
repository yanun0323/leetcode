package main

// HACK: [0017] Golang 2D Slice Copy Issue
var (
	buttons = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	q := [][]byte{}
	for i := range buttons[digits[0]] {
		q = append(q, []byte{buttons[digits[0]][i]})
	}
	for c := 1; c < len(digits); c++ {
		bytes := buttons[digits[c]]
		for count := len(q); count > 0; count-- {
			for j := range bytes {
				buf := make([]byte, len(q[0])+1)
				copy(buf, append(q[0], bytes[j]))
				q = append(q, buf)
			}
			q = q[1:]
		}
	}
	result := make([]string, 0, len(q))
	for i := range q {
		result = append(result, string(q[i]))
	}
	return result
}
