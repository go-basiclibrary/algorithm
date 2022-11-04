package main

// 239
// 单调队列
func maxSlidingWindow02(nums []int, k int) []int {
	p := []int{}

	push := func(i int) {
		for len(p) > 0 && nums[i] >= nums[p[len(p)-1]] {
			p = p[:len(p)-1]
		}
		p = append(p, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, 1, len(nums)-k+1)
	ans[0] = nums[p[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for p[0] <= i-k {
			p = p[1:]
		}
		ans = append(ans, nums[p[0]])
	}

	return ans
}

// 分块+预处理
func maxSlidingWindow03(nums []int, k int) []int {
	prefixMax := make([]int, len(nums))
	suffixMax := make([]int, len(nums))
	for i, v := range nums {
		if i%k == 0 {
			prefixMax[i] = v
		} else {
			prefixMax[i] = max(prefixMax[i-1], v)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || (i+1)%k == 0 {
			suffixMax[i] = nums[i]
		} else {
			suffixMax[i] = max(suffixMax[i+1], nums[i])
		}
	}
	ans := make([]int, len(nums)-k+1)
	for i := range ans {
		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
