package main

// HACK: [0009] Array method -> make a revert number and compare to input argument
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverted := 0
	copy := x
	for copy > 0 {
		reverted = reverted*10 + copy%10
		copy /= 10
	}
	return reverted == x
}

func isPalindrome2(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return reverted == x || reverted/10 == x
}
