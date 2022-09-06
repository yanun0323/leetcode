package main

func searchMatrix(matrix [][]int, target int) bool {
	index := searchFirst(0, matrix, &target)
	return searchRow(matrix[index], &target)
}

func searchFirst(index int, matrix [][]int, target *int) int {
	if len(matrix) == 1 {
		return index
	}

	if matrix[len(matrix)/2][0] > *target {
		return searchFirst(index, matrix[:len(matrix)/2], target)
	}
	return searchFirst(index+len(matrix)/2, matrix[len(matrix)/2:], target)
}

func searchRow(row []int, target *int) bool {
	if len(row) == 1 {
		return row[0] == *target
	}

	if row[len(row)/2] > *target {
		return searchRow(row[:len(row)/2], target)
	}
	return searchRow(row[len(row)/2:], target)
}
