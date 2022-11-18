package main

// 11 盛最多水的容器
// 收敛法,两面假币
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	area := 0
	for r > l {
		area = max(area, min(height[l], height[r])*(r-l)) // 最小高度×宽度
		if height[l] > height[r] {
			r--
		} else {
			l++
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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
