package main

// 11
// 暴力
//func maxArea(height []int) int {
//	max := 0
//	for i := 0; i < len(height)-1; i++ {
//		for j := i + 1; j < len(height); j++ {
//			area := (j - i) * min(height[i], height[j])
//			max = maxF(max, area)
//		}
//	}
//	return max
//}

func maxF(max int, area int) int {
	if max > area {
		return max
	}
	return area
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 两面夹逼
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		max = maxF(max, area)
	}
	return max
}
