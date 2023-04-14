package main

// 32 最长有效括号
// stack 解决方案
func longestValidParentheses(s string) int {
	var maxAns int
	stack := make([]int, 0)
	stack = append(stack, -1) //哨兵

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}

	return maxAns
}

// DP解决方案
func longestValidParentheses02(s string) int {
	var maxAns int
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = 2 + dp[i-2]
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-dp[i-1]-2] + 2 + dp[i-1]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}

	return maxAns
}
