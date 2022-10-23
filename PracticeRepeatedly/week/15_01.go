package main

import "sort"

// 15 三数之和
//暴力

//双指针夹逼
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				r--
				l++
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
