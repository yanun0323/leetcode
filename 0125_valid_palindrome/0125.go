package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	lp := 0
	rp := len(s) - 1
	for {
		findUsefulChar(s, &lp, 1)
		findUsefulChar(s, &rp, -1)
		if lp >= rp {
			return true
		}
		fmt.Println(string(s[lp]), ":", string(s[rp]))
		fmt.Println(s[lp], ":", s[rp])
		fmt.Println(isSameCharacter(s[lp], s[rp]))
		if isSameCharacter(s[lp], s[rp]) {
			lp++
			rp--
			continue
		}
		return false
	}
}
func isSameCharacter(c1, c2 byte) bool {
	if c1 == c2 {
		return true
	}
	if c1 >= 'a' {
		c1 -= 32
	}
	if c2 >= 'a' {
		c2 -= 32
	}
	return c1 == c2
}

func findUsefulChar(s string, p *int, step int) {
	for {
		if isLetter(s[*p]) || edgeOfString(s, p, step > 0) {
			return
		}
		*p += step
	}
}

func isLetter(char byte) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || ('0' <= char && char <= '9')
}

func edgeOfString(s string, index *int, left bool) bool {
	if left {
		return *index == len(s)-1
	}
	return *index == 0
}
