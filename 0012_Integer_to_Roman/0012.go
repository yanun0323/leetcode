package main

var _Hash = map[int]map[int]byte{
	0: map[int]byte{
		1: 'I',
		5: 'V',
	},
	1: map[int]byte{
		1: 'X',
		5: 'L',
	},
	2: map[int]byte{
		1: 'C',
		5: 'D',
	},
	3: map[int]byte{
		1: 'M',
	},
}

func intToRoman(num int) string {
	stack := []byte{}
	for digit := 0; num > 0; digit, num = digit+1, num/10 {
		n := num % 10
		five := n / 5
		one := n % 5

		if one == 4 {
			if five > 0 {
				stack = append(stack, _Hash[digit+1][1])
			} else {
				stack = append(stack, _Hash[digit][5])
			}
			stack = append(stack, _Hash[digit][1])
			continue
		}

		for one > 0 {
			stack = append(stack, _Hash[digit][1])
			one--
		}
		if five > 0 {
			stack = append(stack, _Hash[digit][5])
		}
	}

	result := make([]byte, 0, len(stack))
	for count := len(stack); count > 0; count-- {
		result = append(result, stack[count-1])
	}
	return string(result)
}
