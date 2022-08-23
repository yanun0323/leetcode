package main

func twoSum(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1

	for {

		if target == numbers[lp]+numbers[rp] {
			return []int{lp + 1, rp + 1}
		}

		if target > numbers[lp]+numbers[rp] {
			lp++
			continue
		}
		if target < numbers[lp]+numbers[rp] {
			rp--
			continue
		}
	}
}
