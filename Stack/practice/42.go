package main

// 42 接雨水
func trap(height []int) int {
	// 求最大高度
	heiG := maxHeight(height)
	area := 0
	// 遍历每一行,直到最大高度
	for i := 1; i <= heiG; i++ {
		isStart := false
		tem_Sum := 0
		for j := 0; j < len(height); j++ {
			if isStart && height[j] < i { //开始标记了,并且再次遇到了 hei >= i
				tem_Sum++
			}
			if height[j] >= i {
				area += tem_Sum
				tem_Sum = 0
				isStart = true
			}
		}
	}
	// return
	return area
}

func maxHeight(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	return max
}
