package main

// 42 接雨水
// 横排求
func trap(height []int) int {
	area := 0 //累计面积
	//求出最大高度
	maxH := getMax(height)
	//求每个高度内可接的雨水
	for i := 1; i <= maxH; i++ {
		isStart := false //标记是否开始更新temp
		temp := 0
		for j := 0; j < len(height); j++ {
			if isStart && height[j] < i {
				temp++
			}
			if height[j] >= i {
				area += temp
				temp = 0
				isStart = true
			}
		}
	}
	return area
}

func getMax(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}

	return max
}
