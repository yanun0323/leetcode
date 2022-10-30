package main

func makeIntegerBeautiful(n int64, target int) int64 {
	sum := sumN(n)
	t := int64(target)
	if sum <= t {
		return 0
	}
	var result, ten, fixed int64 = 0, 1, 0
	for sum > t {
		digit := n%10 + fixed
		sum = sum - digit + 1
		result = result + (10-digit)*ten
		ten *= 10
		n /= 10
		fixed = 1
	}
	return result
}

func sumN(n int64) int64 {
	result := int64(0)
	for n > 0 {
		result += n % 10
		n /= 10
	}
	return result
}
