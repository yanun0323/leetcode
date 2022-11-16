package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(int) int { return 0 }

func guessNumber(n int) int {
	l, num := 1, 0
	for l < n {
		num = (l + n) / 2
		switch guess(num) {
		case -1:
			n = num - 1
		case 1:
			l = num + 1
		default:
			return num
		}
	}
	return l
}
