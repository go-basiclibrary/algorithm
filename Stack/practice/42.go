package main

// 42 接雨水
//func trap(height []int) int {
//	// 求最大高度
//	heiG := maxHeight(height)
//	area := 0
//	// 遍历每一行,直到最大高度
//	for i := 1; i <= heiG; i++ {
//		isStart := false
//		tem_Sum := 0
//		for j := 0; j < len(height); j++ {
//			if isStart && height[j] < i { //开始标记了,并且再次遇到了 hei >= i
//				tem_Sum++
//			}
//			if height[j] >= i {
//				area += tem_Sum
//				tem_Sum = 0
//				isStart = true
//			}
//		}
//	}
//	// return
//	return area
//}
//
//func maxHeight(height []int) int {
//	max := 0
//	for i := 0; i < len(height); i++ {
//		if height[i] > max {
//			max = height[i]
//		}
//	}
//	return max
//}

// 42 接雨水
// 动态规划
//func trap(height []int) int {
//	area := 0
//	max_left := make([]int, len(height))
//	max_right := make([]int, len(height))
//	for i := 1; i < len(height); i++ {
//		max_left[i] = maxH(max_left[i-1], height[i-1])
//	}
//	for i := len(height) - 2; i >= 0; i-- {
//		max_right[i] = maxH(max_right[i+1], height[i+1])
//	}
//	for i := 1; i < len(height)-1; i++ {
//		min := minH(max_left[i], max_right[i])
//		if min > height[i] {
//			area += min - height[i]
//		}
//	}
//
//	return area
//}

//func minH(i int, i2 int) int {
//	if i > i2 {
//		return i2
//	}
//	return i
//}

func maxH(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

// 双指针
func trap(height []int) int {
	area := 0
	max_left := 0
	max_right := 0
	left := 1
	right := len(height) - 2
	for i := 1; i < len(height)-1; i++ {
		if height[left-1] < height[right+1] {
			max_left = maxH(max_left, height[left-1])
			min := max_left
			if min > height[left] {
				area += min - height[left]
			}
			left++
		} else {
			max_right = maxH(max_right, height[right+1])
			min := max_right
			if min > height[right] {
				area += min - height[right]
			}
			right--
		}
	}
	return area
}
