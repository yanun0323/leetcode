package main

func majorityElement(nums []int) int {
	n := len(nums) / 2
	hash := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
		if hash[nums[i]] > n {
			return nums[i]
		}
	}
	return 0
}

func BoyerMooreMajorityVoteAlgorithm(nums []int) int {
	count := 0
	major := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
		}
		if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}
