package main

func main() {
	largestRectangleArea([]int{2, 1, 2})
}

// 84 柱状图最大的面积
// 枚举宽度
//func largestRectangleArea(heights []int) int {
//	max := 0
//	for i := 0; i < len(heights); i++ {
//		for j := i; j < len(heights); j++ {
//			minH := math.MaxInt
//			for k := i; k <= j; k++ {
//				minH = minF(heights[k], minH)
//			}
//			area := (j - i + 1) * minH
//			max = maxFunc(area, max)
//		}
//	}
//	return max
//}

func maxFunc(area int, max int) int {
	if area > max {
		return area
	}
	return max
}

func minF(i int, m int) int {
	if i > m {
		return m
	}
	return i
}

// 枚举高度
func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		l, r := i, i
		for l-1 >= 0 && heights[l-1] >= heights[i] {
			l--
		}
		for r+1 < len(heights) && heights[r+1] >= heights[i] {
			r++
		}
		area := (r - l + 1) * heights[i]
		max = maxFunc(area, max)
	}
	return max
}

// 栈
func largestRectangleArea02(heights []int) int {
	max := 0
	stack := make([]int, 0)
	stack = append(stack, -1)    // 哨兵
	heights = append(heights, 0) // 哨兵
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//左边界索引
			l := stack[len(stack)-1]
			area := (i - l - 1) * heights[top]
			max = maxFunc(area, max)
		}
		stack = append(stack, i)
	}
	return max
}
