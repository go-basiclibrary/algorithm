package main

import "fmt"

// 11盛最多水的容器
func main() {
	fmt.Println(maxArea02([]int{1, 1}))
}

/*
	方法1:
		暴力破解法,两个for循环 时间复杂度 O(n^2)  ,超过时间限制,笑死
	方法2:
		收敛法  左右夹逼
*/
// 方法1
func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			max = maxF(max, area)
		}
	}
	return max
}

func maxF(max int, area int) int {
	if max > area {
		return max
	}
	return area
}

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}

// 方法2  收敛法
func maxArea02(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		max = maxF(max, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
