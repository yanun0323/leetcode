package twosum

// AC: true, WA:2 , time: O(n), space: O(n)
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if index, ok := hash[nums[i]]; ok {
			return []int{index, i}
		}
		wanted := target - nums[i]
		hash[wanted] = i
	}
	return nil
}
