package main

func climbStairs(n int) int {
	f := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		if i < 2 {
			f = append(f, 1)
			continue
		}
		f = append(f, f[i-1]+f[i-2])
	}
	return f[n]
}
