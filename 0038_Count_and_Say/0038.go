package main

func countAndSay(n int) string {
	q := []byte{'1'}
	n--
	for ; n > 0; n-- {
		current := byte(' ')
		count := 0
		for i := len(q); i > 0; i-- {
			if current != q[0] && count > 0 {
				q = append(q, byte(count+int('0')), current)
				count = 0
			}

			current = q[0]
			count++
			q = q[1:]
		}
		q = append(q, byte(count+int('0')), current)
	}
	return string(q)
}
