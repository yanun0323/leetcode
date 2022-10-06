package main

import "strings"

func longestCommonPrefix(strs []string) string {
	b := strings.Builder{}
	for char := 0; char < len(strs[0]); char++ {
		for i := 1; i < len(strs); i++ {
			if char == len(strs[i]) || strs[i][char] != strs[i-1][char] {
				return b.String()
			}
		}
		b.WriteByte(strs[0][char])
	}
	return b.String()
}
