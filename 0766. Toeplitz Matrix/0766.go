package main

// FIXME: [0766] Timeout Passed
func isToeplitzMatrix(matrix [][]int) bool {
	for r := range matrix {
		for c := range matrix[r] {
			if r > 0 && c > 0 && matrix[r][c] != matrix[r-1][c-1] {
				return false
			}
		}
	}
	return true
}
