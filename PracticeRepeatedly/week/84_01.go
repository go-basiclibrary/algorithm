package main

// 84 柱状图最大的矩形
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	heights = append(heights, 0) //两个哨兵
	max := 0
	//枚举横坐标
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]
			// 左边界
			l := stack[len(stack)-1]
			//出栈计算面积
			area := (i - l - 1) * heights[top]
			// 计算最大面积
			max = maxFunc(area, max)
		}
		//入栈
		stack = append(stack, i)
	}

	return max
}

func maxFunc(area int, max int) int {
	if area > max {
		return area
	}

	return max
}
