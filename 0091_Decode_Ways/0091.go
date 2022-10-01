package main

func numDecodings(s string) int {
	seen := map[string]int{}
	var decode func(string) int
	decode = func(s string) int {
		if n, exist := seen[s]; exist {
			return n
		}

		if len(s) <= 1 && s != "0" {
			return 1
		}
		if s[0] == '0' {
			return 0
		}

		num := decode(s[1:])
		if valid(s) {
			num += decode(s[2:])
		}
		seen[s] = num
		return num
	}
	return decode(s)
}

func valid(s string) bool {
	return len(s) >= 2 && (s[0]-'0')*10+(s[1]-'0') <= 26
}

func numDecodings2(s string) int {
	if s == "0" {
		return 0
	}
	if len(s) <= 1 {
		return 1
	}
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	if s[0] == '0' {
		dp[1] = 0
	}
	for i := 2; i <= len(s); i++ {
		if valid2[s[i-1:i]] {
			dp[i] = dp[i-1]
		}
		if valid2[s[i-2:i]] {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

var valid2 = map[string]bool{
	"1": true, "2": true, "3": true, "4": true, "5": true,
	"6": true, "7": true, "8": true, "9": true, "10": true,
	"11": true, "12": true, "13": true, "14": true, "15": true,
	"16": true, "17": true, "18": true, "19": true, "20": true,
	"21": true, "22": true, "23": true, "24": true, "25": true,
	"26": true,
}
