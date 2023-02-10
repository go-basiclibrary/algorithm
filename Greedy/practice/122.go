package main

// 122 买卖股票的最佳时机 Ⅱ
// 贪心算法
func maxProfit(prices []int) int {
	var maxPrice int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxPrice += prices[i] - prices[i-1]
		}
	}
	return maxPrice
}
