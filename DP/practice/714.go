package main

// 714 买卖股票的最佳时机含手续费
func maxProfit07(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	var buy, sell = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy, sell = max(buy, sell-prices[i]), max(sell, buy+prices[i]-fee)
	}
	return sell
}
