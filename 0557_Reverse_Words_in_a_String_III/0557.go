package main

import "strings"

func reverseWords2(s string) string {
	b := strings.Builder{}
	stack := []byte{}
	for i := range s {
		if s[i] != ' ' {
			stack = append(stack, s[i])
			continue
		}
		for len(stack) > 0 {
			b.WriteByte(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		b.WriteByte(' ')
	}

	for len(stack) > 0 {
		b.WriteByte(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return b.String()
}

func reverseWords(s string) string {
	b := strings.Builder{}
	slow, fast := 0, 0
	for ; len(s) > fast; fast++ {
		if s[fast] != ' ' {
			continue
		}
		for i := fast - 1; i >= slow; i-- {
			b.WriteByte(s[i])
		}
		b.WriteByte(' ')
		slow = fast + 1
	}
	for i := fast - 1; i >= slow; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
