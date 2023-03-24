package main

// 188 买卖股票的最佳时机IV
func maxProfit06(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var buyj, sellj = make([]int, k+1), make([]int, k+1)
	for i := range buyj {
		buyj[i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buyj[j] = max(buyj[j], sellj[j-1]-prices[i])
			sellj[j] = max(sellj[j], buyj[j]+prices[i])
		}
	}
	return sellj[k]
}
