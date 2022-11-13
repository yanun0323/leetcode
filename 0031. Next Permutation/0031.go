package main

func nextPermutation(nums []int) {
	currentIndex := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			currentIndex = i
			break
		}
	}
	for i := currentIndex + 2; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[currentIndex] {
			nums[currentIndex], nums[i] = nums[i], nums[currentIndex]
			break
		}
	}
	if currentIndex >= 0 {
		reverse(currentIndex+1, nums)
		return
	}
	reverse(0, nums)
}

func reverse(l int, nums []int) {
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func nextPermutation_Old(nums []int) {
	minIndex := -1
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		minIndex = -1
		for j := i + 1; j < length; j++ {
			if nums[j] > nums[i] && (minIndex < 0 || nums[j] < nums[minIndex]) {
				minIndex = j
			}
		}
		if minIndex >= 0 {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
			i++
			for ; i < length; i++ {
				minIndex = i
				for j := i + 1; j < length; j++ {
					if nums[j] < nums[minIndex] {
						minIndex = j
					}
				}
				nums[i], nums[minIndex] = nums[minIndex], nums[i]
			}
			return
		}
	}
	l, r := 0, length-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
