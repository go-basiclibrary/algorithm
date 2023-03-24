package main

// 122 买卖股票的最佳时机Ⅱ
func maxProfit02(prices []int) int {
	// 贪心,DP
	var maxPrice int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxPrice += prices[i] - prices[i-1]
		}
	}
	return maxPrice
}
