package main

import "math"

func main() {
	coinChange([]int{1, 2, 5}, 11)
}

func coinChange(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && memo[i] > memo[i-c]+1 {
				memo[i] = memo[i-c] + 1
			}
		}
	}
	if memo[amount] == math.MaxInt32 {
		return -1
	}
	return memo[amount]
}

// DP 方程解决问题
func coinChangeDP(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memo[i] = amount + 1
		for _, c := range coins {
			if i >= c { // 当面额大于硬币价值时
				memo[i] = min(memo[i], memo[i-c]+1) // 找组合中小的那个值
			}
		}
	}
	if memo[amount] == amount+1 {
		return -1
	}
	return memo[amount]
}

func min(i int, i2 int) int {
	if i > i2 {
		return i2
	}
	return i
}
