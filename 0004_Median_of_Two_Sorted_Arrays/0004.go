package main

// 2 4 9 10|16			5
// 8 12|19 20 22 28			6
const (
	_MaxUint = ^uint(0)
	_MaxInt  = int(_MaxUint >> 1)
	_MinInt  = -_MaxInt - 1
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	l := 0
	r := m
	for l <= r {
		p1 := (l + r) >> 1
		p2 := (m+n+1)>>1 - p1
		maxL1, maxL2 := _MinInt, _MinInt
		if p1 > 0 {
			maxL1 = nums1[p1-1]
		}
		if p2 > 0 {
			maxL2 = nums2[p2-1]
		}

		minR1, minR2 := _MaxInt, _MaxInt
		if p1 < m {
			minR1 = nums1[p1]
		}
		if p2 < n {
			minR2 = nums2[p2]
		}

		if maxL1 <= minR2 && maxL2 <= minR1 {
			if (m+n)%2 == 0 {
				return float64(max(maxL1, maxL2)+min(minR1, minR2)) / 2
			}
			return float64(max(maxL1, maxL2))
		}

		if maxL1 > minR2 {
			r = p1 - 1
			continue
		}
		l = p1 + 1

	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	p1 := 0
	p2 := 0
	for p1 < len(nums1) || p2 < len(nums2) {
		if p1 == len(nums1) {
			merged = append(merged, nums2[p2:]...)
			break
		}
		if p2 == len(nums2) {
			merged = append(merged, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
			continue
		}
		merged = append(merged, nums2[p2])
		p2++
		continue
	}

	if len(merged)%2 == 1 {
		return float64(merged[len(merged)/2])
	}

	return (float64(merged[len(merged)/2]) + float64(merged[len(merged)/2-1])) / 2

}
