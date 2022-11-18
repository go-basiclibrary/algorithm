package main

// 42接雨水,动态规划
func trap02(height []int) int {
	area := 0
	max_left := make([]int, len(height))
	max_right := make([]int, len(height))
	// 求i左边墙的高度
	for i := 1; i < len(height); i++ {
		max_left[i] = maxH(max_left[i-1], height[i-1])
	}
	// 求i右边墙的比较高的
	for i := len(height) - 2; i >= 0; i-- {
		max_right[i] = maxH(max_right[i+1], height[i+1])
	}
	// 求左右两边矮的高度,如果大于当前要求的i的高度,那么则可以计算面积(可以装水)
	// 要考虑两边的最边界
	for i := 1; i < len(height)-1; i++ {
		h := minH(max_left[i], max_right[i])
		if h > height[i] {
			area += h - height[i]
		}
	}
	return area
}

func minH(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxH(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
