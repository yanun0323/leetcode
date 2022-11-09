package main

func maximum69Number(num int) int {
	n := num
	ten := 1
	found := -1
	for n > 0 {
		if n%10 == 6 {
			found = ten
		}
		n /= 10
		ten *= 10
	}
	if found == -1 {
		return num
	}

	return num + ten*3
}
