package main

import "math"

// BUG: [0743] 22.11.08 Stuck
type Transport struct {
	Target int
	Time   int
}

/* Dijkstra */
func networkDelayTime(times [][]int, n int, k int) int {
	m := map[int][]Transport{}
	for _, time := range times {
		m[time[0]] = append(m[time[0]], Transport{
			Target: time[1],
			Time:   time[2],
		})
	}
	result := 0
	g := make([]int, n+1)
	var dijkstra func(int)
	dijkstra = func(cur int) {
		now := g[cur]
		g[cur] = -1
		next := -1
		min := math.MaxInt
		for _, ts := range m[cur] {
			if g[ts.Target] == -1 {
				continue
			}
			time := now + ts.Time
			if time < g[ts.Target] {
				g[ts.Target] = time
			} else {
				time = g[ts.Target]
			}
			if time < min {
				next = ts.Target
				min = time
			}
		}
		if next != -1 {
			dijkstra(next)
			return
		}
		if now > result {
			result = now
		}
	}
	dijkstra(k)
	return result
}

/* Origin */
func networkDelayTime2(times [][]int, n int, k int) int {
	m := make([][]Transport, n+1)
	for _, t := range times {
		m[t[0]] = append(m[t[0]], Transport{
			Target: t[1],
			Time:   t[2],
		})
	}

	g := make([]int, n+1)
	for i := range g {
		g[i] = -1
	}
	now := 0
	var dfs func(n int)
	dfs = func(n int) {
		if g[n] > -1 && now > g[n] {
			return
		}
		g[n] = now
		for _, ts := range m[n] {
			now += ts.Time
			dfs(ts.Target)
			now -= ts.Time
		}
	}
	g[k] = 0
	dfs(k)
	result := 0
	for i := 1; i < len(g); i++ {
		if g[i] == -1 {
			return -1
		}
		if g[i] > result {
			result = g[i]
		}
	}
	return result
}
