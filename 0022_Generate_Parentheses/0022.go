package main

func generateParenthesis(n int) []string {
	return Get(make([]string, 0, (n+1)*n/2), make([]byte, 0, n*2), n, n)
}

func Get(answer []string, buf []byte, l, r int) []string {
	if l+r == 0 {
		return append(answer, string(buf))
	}
	if l < r {
		answer = Get(answer, append(buf, ')'), l, r-1)
	}
	if l > 0 {
		answer = Get(answer, append(buf, '('), l-1, r)
	}
	return answer
}
