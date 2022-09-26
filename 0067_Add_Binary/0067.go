package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	add := false
	stack := make([]byte, len(a)+1)
	top := len(a)
	for countA, countB := len(a)-1, len(b)-1; countA >= 0; countA, countB = countA-1, countB-1 {
		digit := 0
		if countB >= 0 && b[countB] == '1' {
			digit++
		}
		if a[countA] == '1' {
			digit++
		}

		if add {
			digit++
			add = false
		}

		switch digit {
		case 0:
			stack[top] = '0'
			top--
		case 1:
			stack[top] = '1'
			top--
		case 2:
			stack[top] = '0'
			top--
			add = true
		case 3:
			stack[top] = '1'
			top--
			add = true
		}
	}
	if add {
		stack[top] = '1'
		return string(stack)
	}

	return string(stack[1:])
}
