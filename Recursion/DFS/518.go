package main

func main() {
	change(5, []int{1, 2, 5})
}

// dfs 解决
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
