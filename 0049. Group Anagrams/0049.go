package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	table := map[[26]int]int{}
	result := [][]string{}
	for i := range strs {
		k := hash(strs[i])
		if idx, exist := table[k]; exist {
			result[idx] = append(result[idx], strs[i])
			continue
		}
		result = append(result, []string{strs[i]})
		table[k] = len(result) - 1
	}
	return result
}

func hash(s string) [26]int {
	result := [26]int{}
	for i := range s {
		result[int(s[i]-'a')]++
	}
	return result
}

func groupAnagrams2(strs []string) [][]string {
	table := map[string]int{}
	result := [][]string{}
	for i := range strs {
		sli := []byte(strs[i])
		sort.Slice(sli, func(f, s int) bool {
			return sli[f] < sli[s]
		})
		s := string(sli)
		if idx, exist := table[s]; exist {
			result[idx] = append(result[idx], strs[i])
			continue
		}
		result = append(result, []string{strs[i]})
		table[s] = len(result) - 1
	}
	return result
}
