package main

func numDecodings(s string) int {
	var dp = make([]int, len(s)+1)

	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

// 优化空间
func numDecodings02(s string) int {
	var a, b, c = 0, 1, 0

	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			c += a
		}
		a, b = b, c
	}
	return c
}
