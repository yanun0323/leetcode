package main

// FIXME: [1155]
// DP with array as memory
// example:
// diceCount = 3, maxDiceNum = 5, targetNum = 10
// first dice could be 1 to 5.
// therefore, the left 2 dice could be the same case as:
// [diceCount = 2, maxDiceNum = 5, targetNum = 5 to 9 (10 sub by first dice)]
//
// We can create a 2D array that store numRolls of all (diceCount vs targetNumber)
// like: numRolls [][]int => [diceCount][targetNum]int
//
// diceCount = 3, maxDiceNum = 5, targetNum = 10 could be :
// numRolls[3][10] = numRolls[2][10-1] + numRolls[2][10-2] + ... + numRolls[2][10-5]

func numRollsToTarget(n int, k int, target int) int {
	numRolls := make([][]int, n+1)
	for i := range numRolls {
		numRolls[i] = make([]int, target+1)
	}

	for num := 1; num <= k && num <= target; num++ {
		numRolls[1][num] = 1
	}

	mod := 1_000_000_000 + 7
	for diceCount := 2; diceCount <= n; diceCount++ {
		for targetNum := 1; targetNum <= target; targetNum++ {
			for firstDiceNum := 1; firstDiceNum <= k && firstDiceNum < targetNum; firstDiceNum++ {
				numRolls[diceCount][targetNum] += numRolls[diceCount-1][targetNum-firstDiceNum]
				numRolls[diceCount][targetNum] %= mod
			}
		}
	}
	return numRolls[n][target]
}
