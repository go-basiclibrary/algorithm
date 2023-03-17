package main

import "math"

// 121 买卖股票的最佳时机
func maxProfit(prices []int) int {
	var minPri, maxPri int
	minPri = math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPri {
			minPri = prices[i]
		} else if prices[i]-minPri > maxPri {
			maxPri = prices[i] - minPri
		}
	}
	return maxPri
}
