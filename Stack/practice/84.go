package main

import (
	"fmt"
	"math"
)

//84 柱状图最大的矩形
func main() {
	fmt.Println(largestRectangleArea05([]int{7, 7, 1}))
}

/*
	三种方法:
		1.横向枚举
		2.纵向枚举
		3.栈
*/
//1.枚举宽度
func largestRectangleArea(heights []int) int {
	max := 0
	//双向枚举,找范围内最小高度
	for i := 0; i < len(heights); i++ {
		min := math.MaxInt
		for j := i; j < len(heights); j++ {
			min = minHeight(min, heights[j])
			area := (j - i + 1) * min
			max = maxArea(max, area)
		}
	}
	return max
}

// 2.枚举高度
func largestRectangleArea02(heights []int) int {
	max := 0
	//双向枚举,找范围内最小宽度
	for i := 0; i < len(heights); i++ {
		l, r := i, i
		h := heights[i]
		for l-1 >= 0 && heights[l-1] >= h {
			l--
		}
		for r+1 < len(heights) && heights[r+1] >= h {
			r++
		}
		area := (r - l + 1) * h
		max = maxArea(max, area)
	}
	return max
}

func maxArea(max int, area int) int {
	if max > area {
		return max
	}
	return area
}

func minHeight(i int, i2 int) int {
	if i > i2 {
		return i2
	}

	return i
}

// 3.栈
func largestRectangleArea03(heights []int) int {
	max := 0
	stack := make([]int, 0)
	stack = append(stack, -1)    //控制左边界
	heights = append(heights, 0) //引入一个哨兵
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			//拿出栈顶元素的索引
			top := stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]
			// 左边界索引
			l := stack[len(stack)-1]

			area := (i - l - 1) * heights[top]
			max = maxArea(area, max)
		}
		stack = append(stack, i)
	}

	return max
}

func largestRectangleArea0302(heights []int) int {
	max := 0
	stack := make([]int, 0)
	heights = append(heights, 0)        //引入一个哨兵
	for i := 0; i < len(heights); i++ { //右边界
		for len(stack) > 0 && heights[len(stack)-1] > heights[i] {
			// 元素
			top := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			//计算右边界
			l := 0
			if len(stack) > 0 {
				l = stack[len(stack)-1]
			}
			area := (i - l - 1) * heights[top]
			max = maxArea(area, max)
		}
		stack = append(stack, i)
	}

	return max
}

func largestRectangleArea04(hs []int) int {
	var ms []int        //1)单调栈，保存的是index
	hs = append(hs, -1) //2)哨兵
	var max int
	for i, h := range hs {
		for len(ms) > 0 && hs[ms[len(ms)-1]] > h {
			index := -1
			if len(ms) >= 2 { //4)栈里面倒数的值，已经处理了
				index = ms[len(ms)-2]
			}
			height := hs[ms[len(ms)-1]]
			width := i - index - 1
			if area := height * width; max < area {
				max = area
			}
			ms = ms[:len(ms)-1] //3)完成出栈
		}
		ms = append(ms, i)
	}
	return max
}

func largestRectangleArea05(heights []int) int {
	//单调栈（单调递增）
	stack := make([]int, 0)
	stack = append(stack, -1)    //stack的哨兵，方便确定左边界
	heights = append(heights, 0) //添加一个哨兵，减少代码量
	ln := len(heights)
	res := 0 //结果

	for i := 0; i < ln; i++ {
		//因为我们无法访问heights[-1]，所以限制len(stack) > 1
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			//栈顶元素，也就是当前要求的矩形柱子的下标
			top := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//左边界（栈顶元素的后一个元素）
			l := stack[len(stack)-1]
			//矩形面积：(右边界-左边界-1) * 高度
			//右边界就是i
			//高度就是以栈顶元素为下标的柱子的高度
			//左边界就是栈顶元素的下一位元素（因为我们添加了哨兵-1，所以这公式依旧成立）
			res = maxArea(res, (i-l-1)*heights[top])
		}
		stack = append(stack, i)

	}

	return res
}
