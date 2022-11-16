package main

import "bytes"

// BUG: [0005] 22.11.15 Stuck
func longestPalindrome(s string) string {
	var buffer bytes.Buffer
	buffer.WriteByte('*')
	for i := 0; i < len(s); i++ {
		buffer.WriteByte(s[i])
		buffer.WriteByte('*')
	}
	sp := buffer.String()
	dp := make([]int, len(sp))
	center := 0
	radius := 0
	for center < len(sp) {
		for center-radius-1 >= 0 && center+radius+1 < len(sp) {
			if sp[center-radius-1] != sp[center+radius+1] {
				break
			}
			radius++
		}
		dp[center] = radius
		center++
		radius = 0
	}
	max := 0
	maxIdx := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
			maxIdx = i
		}
	}
	ss := sp[maxIdx-max : maxIdx+max+1]
	buffer.Reset()
	for i := 1; i < len(ss); i += 2 {
		buffer.WriteByte(ss[i])
	}
	return buffer.String()
}
