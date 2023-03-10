package main

// 152 乘积最大数组
func maxProduct(nums []int) int {
	maxF, minF, max := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		xF, mF := maxF, minF
		maxF = maxWW(xF*nums[i], maxWW(nums[i], mF*nums[i])) // 内部比较负数
		minF = minWW(mF*nums[i], minWW(nums[i], xF*nums[i]))
		max = maxWW(maxF, max)
	}
	return max
}

func minWW(i int, f int) int {
	if i > f {
		return f
	}
	return i
}

func maxWW(i int, f int) int {
	if i > f {
		return i
	}
	return f
}
