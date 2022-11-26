package main

// 84 柱状图中最大的矩形
// 维护栈
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	area := 0

	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			//用于计算的高度
			top := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//用于计算宽度
			l := stack[len(stack)-1]
			area = maxMj(area, (i-l-1)*heights[top])
		}
		stack = append(stack, i)
	}

	return area
}

func maxMj(a, b int) int {
	if a > b {
		return a
	}
	return b
}
