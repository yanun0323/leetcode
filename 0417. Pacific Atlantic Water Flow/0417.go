package main

import "math"

var (
	_P     int8 = 1 << 1
	_A     int8 = 1 << 2
	_Valid      = _A | _P
)

func pacificAtlantic(heights [][]int) [][]int {
	checked := create2DBitArr(len(heights), len(heights[0]))

	result := [][]int{}
	var lowToHigh func(int8, int, int, int)
	lowToHigh = func(bit int8, prevVal, r, c int) {
		if r < 0 || c < 0 || r >= len(heights) || c >= len(heights[0]) || heights[r][c] < prevVal || checked[r][c]&bit > 0 {
			return
		}
		checked[r][c] = checked[r][c] | bit
		val := heights[r][c]
		heights[r][c] = math.MinInt
		lowToHigh(bit, val, r-1, c)
		lowToHigh(bit, val, r, c-1)
		lowToHigh(bit, val, r+1, c)
		lowToHigh(bit, val, r, c+1)
		heights[r][c] = val
	}

	for i := range heights {
		lowToHigh(_P, 0, i, 0)
		lowToHigh(_A, 0, i, len(heights[0])-1)
	}
	for i := range heights[0] {
		lowToHigh(_P, 0, 0, i)
		lowToHigh(_A, 0, len(heights)-1, i)
	}
	for r := range checked {
		for c := range checked[r] {
			if checked[r][c] == _Valid {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

func create2DBitArr(r, c int) [][]int8 {
	result := make([][]int8, r)
	for i := range result {
		result[i] = make([]int8, c)
	}
	return result
}
