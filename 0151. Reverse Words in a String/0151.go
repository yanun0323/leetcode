package main

import "strings"

func reverseWords(s string) string {
	stack := []string{}
	temp := []byte{}
	for i := range s {
		if s[i] != ' ' {
			temp = append(temp, s[i])
			continue
		}

		if len(temp) == 0 {
			continue
		}

		stack = append(stack, string(temp))
		temp = []byte{}
	}

	if len(temp) > 0 {
		stack = append(stack, string(temp))
	}
	b := strings.Builder{}
	for len(stack) > 0 {
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return b.String()
}
