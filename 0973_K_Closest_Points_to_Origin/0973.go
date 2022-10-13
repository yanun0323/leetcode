package main

import "sort"

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return distance(points[i]) < distance(points[j])
	})
	return points[:k]
}

func distance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
