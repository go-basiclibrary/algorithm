package main

// 42 接雨水
// func1 单调栈
//func trap(height []int) int {
//	var area int
//	var trapStack = make([]int, 0, len(height))
//	for i, h := range height {
//		// 如果当前trapStack中有元素,且当前h大于左边,则进入循环出栈进行计算area
//		for len(trapStack) > 0 && h > height[trapStack[len(trapStack)-1]] {
//			top := trapStack[len(trapStack)-1]
//			// pop top
//			trapStack = trapStack[:len(trapStack)-1]
//			// 如果左侧没有边界,则直接退出
//			if len(trapStack) == 0 {
//				break
//			}
//			// 获取左边界
//			left := trapStack[len(trapStack)-1]
//			// 获取宽度
//			curWeight := i - left - 1
//			// 获取高度
//			curHeight := min(h, height[left]) - height[top]
//			// 计算面积
//			area += curWeight * curHeight
//		}
//		trapStack = append(trapStack, i)
//	}
//	return area
//}
//
//func min(h int, i int) int {
//	if h > i {
//		return i
//	}
//	return h
//}

// func2 双指针用法
//func trap(height []int) int {
//	var area int
//	// 初始化双指针起点
//	var left, right = 0, len(height) - 1
//	// 记录左侧最大值和右侧最大值
//	var leftMax, rightMax = 0, 0
//	for left < right {
//		leftMax = maxFunc(leftMax, height[left])
//		rightMax = maxFunc(rightMax, height[right])
//		// 如果当前left高度,小于right高度
//		if height[left] < height[right] {
//			area += leftMax - height[left]
//			left++
//		} else { // 如果当前left高度,大于等于right高度
//			area += rightMax - height[right]
//			right--
//		}
//	}
//	return area
//}
//
//func maxFunc(leftMax int, h int) int {
//	if leftMax > h {
//		return leftMax
//	} else {
//		return h
//	}
//}
