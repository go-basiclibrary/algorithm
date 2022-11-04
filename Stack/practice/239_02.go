package main

// 239 最大滑动窗口  单调队列
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{} //队列
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] { // i 大于队尾,则清理队尾元素
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, 1, len(nums)-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for q[0] <= i-k {
			// 循环最大元素出栈
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

// 分块+预处理
func maxSlidingWindow02(nums []int, k int) []int {
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

func max(i int, v int) int {
	if i > v {
		return i
	}
	return v
}
