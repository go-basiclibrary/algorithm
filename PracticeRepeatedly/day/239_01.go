package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var a []int

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	fmt.Println("push", h.IntSlice)
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

// 239 滑动窗口最大值
// heap(优先队列)去做
func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	h := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		// 存储index
		h.IntSlice[i] = i
	}

	heap.Init(h)
	ans := make([]int, 1, len(nums)-k+1)
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
