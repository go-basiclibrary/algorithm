package main

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(val interface{}) {
	h.IntSlice = append(h.IntSlice, val.(int))
}

func (h *hp) Pop() interface{} {
	val := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return val
}

func (h *hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

// 239 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	h := &hp{make([]int, 0)}
	for i := 0; i < k; i++ {
		h.IntSlice = append(h.IntSlice, i)
	}
	heap.Init(h)

	ans := make([]int, 1, len(nums)-2)
	ans[0] = nums[h.IntSlice[0]]
	for i := k; i < len(nums); i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[h.IntSlice[0]])
	}
	return ans
}

// 单调栈
func maxSlidingWindow02(nums []int, k int) []int {
	//单调栈,先只存放最大的数
	stack := []int{}
	push := func(i int) {
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int, 1, len(nums)-k+1)
	ans[0] = nums[stack[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for stack[0] <= i-k {
			stack = stack[1:]
		}
		ans = append(ans, nums[stack[0]])
	}
	return ans
}
