package main

func search(nums []int, target int) int {
	rotateIndex := searchRotateIndex(0, len(nums)-1, nums)
	if rotateIndex == 0 || target < nums[0] {
		return searchTarget(rotateIndex, len(nums)-1, nums, &target)
	}
	return searchTarget(0, rotateIndex-1, nums, &target)
}

func searchRotateIndex(l, r int, nums []int) int {
	if l >= r {
		return l
	}

	mid := (l + r) >> 1
	if nums[mid] > nums[r] {
		return searchRotateIndex(mid+1, r, nums)
	}
	return searchRotateIndex(l, mid, nums)
}

func searchTarget(l, r int, nums []int, target *int) int {
	if l >= r {
		if nums[l] == *target {
			return l
		}
		return -1
	}
	mid := (l + r) >> 1
	if *target > nums[mid] {
		return searchTarget(mid+1, r, nums, target)
	}
	return searchTarget(l, mid, nums, target)
}
