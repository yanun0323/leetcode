package main

// nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//   1 2 3 2 1 x   max=3
// 3 0 0 3 0 0 0
// 2 0 1 0 2 0 0
// 1 1 0 0 0 1 0
// 4 0 0 0 0 0 0
// 7 0 0 0 0 0 0
// x 0 0 0 0 0 0

func findLength(nums1 []int, nums2 []int) int {
	max := 0
	store := make([][]int, 0, len(nums1)+1)
	for i := range store {
		store[i] = make([]int, 0, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i] == nums2[j] {
				store[i][j] = 1 + store[i-1][j-1]
				if max < store[i][j] {
					max = store[i][j]
				}
			}
		}
	}
	return max
}
