package main

func groupAnagrams_old(strs []string) [][]string {
	letterA := int('a')
	hash := make(map[uint64][]string, 0)
	var key uint64
	for i := 0; i < len(strs); i++ {
		key = 0
		for j := 0; j < len(strs[i]); j++ {
			key += createKey(int(strs[i][j]) - letterA)
		}
		hash[key] = append(hash[key], strs[i])
	}
	var result [][]string
	for _, v := range hash {
		result = append(result, v)
	}
	return result
}

func createKey(time int) uint64 {
	var num uint64 = 1
	for i := 0; i < time; i++ {
		num *= 101
	}
	return num
}
