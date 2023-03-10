package main

// 198 打家劫舍  DP + array
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		res[i] = max(res[i-1], res[i-2]+nums[i])
	}

	return res[len(nums)-1]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

// 优化 滚动数组
func rob02(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}

	return second
}
