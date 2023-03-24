package main

//  最佳买卖股票时机含冷冻期
func maxProfit04(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// f[i][0]: 手上持有股票且收益最大
	// f[i][1]: 手上不持有股票,并且处于冷冻期,无法购买股票
	// f[i][2]: 手上不持有股票,且不处于冷冻期,可以购买股票
	f := make([][3]int, len(prices))
	f[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// 目前持有股票,要么i-1天已经持有   f[i-1][0]
		// 或者第i天买入,那么i-1天就不能持有股票并且不是冷冻期
		// f[i-1][2]-prices[i]  今天买入
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		// f[i][1] 第i天结束之后处于冷冻期,说明i-1天一定进行了卖股票
		f[i][1] = f[i-1][0] + prices[i]
		// f[i][2],说明我们第i天结束之后,不持有任何股票并且不处于冷冻期,说明当天没有进行任何操作
		// 如果处于冷冻f[i-1][1],不处于冷冻期
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[len(prices)-1][1], f[len(prices)-1][2])
}

// 优化空间
func maxProfit05(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var f [3]int
	f[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		f[0], f[1], f[2] = max(f[0], f[2]-prices[i]), f[0]+prices[i], max(f[1], f[2])
	}
	return max(f[1], f[2])
}
