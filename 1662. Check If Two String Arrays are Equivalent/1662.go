package main

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	c1 := strings.Join(word1, "")
	c2 := strings.Join(word2, "")
	return c1 == c2
}
