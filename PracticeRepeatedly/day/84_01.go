package main

import (
	"fmt"
	"math"
)

// 84 柱状图最大的矩形
func main() {
	fmt.Println(largestRectangleArea02([]int{2, 1, 5, 6, 2, 3}))
}

/*
	三种方法:
		1.暴力for,枚举宽度
		2.暴力for,枚举高度
		3.栈,遇到小高度,确定右边界
*/
func largestRectangleArea(heights []int) int {
	l := 0
	r := 0

	max := 0
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			minHeight := math.MaxInt
			l, r = i, j
			for l <= r {
				minHeight = minH(minHeight, heights[l])
				l++
			}
			area := minHeight * (j - i + 1)
			max = maxF84(area, max)
		}
	}
	return max
}
func maxF84(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minH(height int, i int) int {
	if height > i {
		return i
	}
	return height
}

// 枚举高度
func largestRectangleArea01(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		l, r := i, i
		h := heights[i]
		for l-1 >= 0 && h <= heights[l-1] {
			l--
		}
		for r+1 < len(heights) && h <= heights[r+1] {
			r++
		}
		area := (r - l + 1) * h
		max = maxF84(area, max)
	}

	return max
}

// 栈
func largestRectangleArea02(heights []int) int {
	max := 0
	stack := make([]int, 0)
	stack = append(stack, -1)    //引入哨兵  索引接入-1
	heights = append(heights, 0) // 引入哨兵,最后将所有栈中的元素pop出去
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			area := (i - l - 1) * heights[top]
			max = maxF84(area, max)
		}
		stack = append(stack, i)
	}
	return max
}
