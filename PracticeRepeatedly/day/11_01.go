package main

// 11 盛最多水的容器
func main() {

}

// 方法一:暴力解决
func maxArea(height []int) int {
	max := 0
	//双层for循环
	for i := 0; i < len(height); i++ {
		for j := 1; j < len(height); j++ {
			//横纵相乘,暴力解法获取max
			area := (j - i) * min(height[j], height[i])
			max = maxF(area, max)
		}
	}

	return max
}

// return max area
func maxF(area int, max int) int {
	if area > max {
		return area
	}
	return max
}

// 两者只能使用小高度才能盛下水
func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}

// 方法二:收敛法,两边夹逼
func maxArea02(height []int) int {
	max := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		max = maxF(area, max)
		if height[left] > height[right] { //较短的高度收敛
			right--
		} else {
			left++
		}
	}

	return max
}
