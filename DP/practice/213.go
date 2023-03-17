package main

// 213 打家劫舍Ⅱ
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// 两种情况,要么偷 0,n-2家  那么偷1,n-1家
	return max(dp(nums, 0, len(nums)-2), dp(nums, 1, len(nums)-1))
}

func dp(nums []int, x, y int) int {
	first := nums[x]
	second := max(nums[x], nums[x+1])
	for i := x + 2; i < y; i++ {
		first, second = second, max(second, first+nums[i])
	}
	return second
}
