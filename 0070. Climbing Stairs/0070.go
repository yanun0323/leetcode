package main

func climbStairs(n int) int {
	dp := []int{1, 1}
	for n > 1 {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
		n--
	}
	return dp[1]
}
