package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testCase := map[int][]int{
		49: {1, 8, 6, 2, 5, 4, 8, 3, 7},
	}

	for answer, height := range testCase {
		assert.Equal(t, answer, maxArea(height))
	}
}
