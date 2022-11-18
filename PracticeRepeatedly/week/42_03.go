package main

// 42 接雨水
// 双指针求法
func trap03(height []int) int {
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[leftMax] < height[rightMax] {
			//执行面积
			area += leftMax - height[left]
			left++
		} else {
			area += rightMax - height[right]
			right++
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
