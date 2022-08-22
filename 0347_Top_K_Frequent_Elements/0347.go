package main

// WA, time: O(n), space: O(n)
func topKFrequent(nums []int, k int) []int {
	frq := make(map[int][]int, len(nums)) // [次數]數字
	seen := make(map[int]int, len(nums))  // [數字] 次數
	for i := 0; i < len(nums); i++ {
		seen[nums[i]]++
	}

	for k, v := range seen {
		frq[v] = append(frq[v], k)
	}
	result := make([]int, 0, k)
	count := len(nums)
	for {
		if len(frq[count]) > 0 {
			result = append(result, frq[count]...)
			k -= len(frq[count])
			if k <= 0 {
				return result
			}
		}
		count--
	}
}
