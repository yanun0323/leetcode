package main

func validUtf82(data []int) bool {
	skip := 0
	var b byte
	for i := range data {
		b = ^byte(data[i])
		if skip > 0 {
			if b>>6 != 1 {
				return false
			}
			skip--
			continue
		}
		b = b >> 3
		// 4, 3, 2
		for n := 0; n < 3; n++ {
			if b == 1 {
				skip = 3 - n
				break
			}
			b = b >> 1
		}
		// 1
		if skip == 0 && b>>1 != 1 {
			return false
		}
	}
	return skip == 0
}

func validUtf8(data []int) bool {
	tail := 0
	for i := range data {
		if tail > 0 {
			if data[i]>>6 != 2 {
				return false
			}
			tail--
			continue
		}
		mask := 1 << 7
		for data[i]&mask != 0 {
			tail++
			mask = mask >> 1
		}
		if tail == 0 {
			continue
		}
		if tail == 1 || tail > 4 {
			return false
		}
		tail--
	}
	return tail == 0
}
