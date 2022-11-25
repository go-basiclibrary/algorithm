package main

import "sort"

// 26
// 获取数组中重复项的元素
func removeDuplicates(nums []int) int {
	var diff [10]int
	for i := 0; i < len(nums); i++ {
		diff[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := 0
	for _, v := range diff {
		if v > 0 {
			res++
		}
	}
	return res
}

func removeDuplicates02(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	j++
	return j
}
