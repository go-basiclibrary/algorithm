package main

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct {
	sort.IntSlice //实现优先队列,这里是内嵌加重载
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1] // 包左不包右
	return v
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

// 239 滑动窗口最大值
// 方法一,排序堆+切片
func maxSlidingWindow01(nums []int, k int) []int {
	a = nums
	h := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}

	heap.Init(h)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[h.IntSlice[0]]

	for i := k; i < n; i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[h.IntSlice[0]])
	}

	return ans
}
