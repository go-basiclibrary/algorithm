package main

// 买卖股票的最佳时机Ⅲ
func maxProfit03(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 未进行任何操作
		// 只进行一次买入操作
		buy1 = max(buy1, -prices[i]) // 进行比较负数,看谁更大,则谁更小
		// 进行了一次买操作和一次卖操作,即完成了一笔交易
		sell1 = max(sell1, prices[i]+buy1)
		// 在完成了一笔交易的前提下,进行了第二次买操作
		buy2 = max(buy2, sell1-prices[i])
		// 完成全部两笔交易
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
