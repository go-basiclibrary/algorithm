package main

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}

// 42 接雨水
// 横向求面积
func trap(height []int) int {
	area := 0
	//先求出最大高度
	mh := maxHeight(height)
	//依次从1到最高高度遍历横排
	for i := 1; i <= mh; i++ {
		isSign := false // 开始标记
		temp := 0       // 用于计算单横向面积
		//循环遍历高度
		for j := 0; j < len(height); j++ {
			if isSign && height[j] < i {
				temp++ //求出单元面积
			}
			if height[j] >= i { // 当再次遇到高度才能形成池
				//收尾
				area += temp
				temp = 0
				isSign = true
			}
		}
	}
	return area
}

func maxHeight(h []int) int {
	max := 0
	for _, v := range h {
		if v > max {
			max = v
		}
	}
	return max
}

// 动态规划
func trap02(height []int) int {
	area := 0
	maxL := make([]int, len(height))
	maxR := make([]int, len(height))
	// 先求出每一列i对应的最左边的最高高度
	for i := 1; i < len(height); i++ {
		maxL[i] = maxFunc(maxL[i-1], height[i-1])
	}
	// 再求出每一列i对应的最右边的最高高度
	for i := len(height) - 2; i >= 0; i-- {
		maxR[i] = maxFunc(maxR[i+1], height[i+1])
	}
	// 根据每一列i进行比对较低高度,求高度差即是面积
	for i := 0; i < len(height); i++ {
		h := minFunc(maxL[i], maxR[i])
		if h > height[i] {
			area += h - height[i]
		}
	}
	return area
}

func minFunc(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxFunc(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
