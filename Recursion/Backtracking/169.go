package main

import "sort"

func main() {
	//majorityElement([]int{3, 2, 3})
}

// 169 多数元素

//func majorityElement(nums []int) int {
//	var n = len(nums)
//	if n == 1 {
//		return nums[0]
//	}
//
//	var m = make(map[int]int)
//	for i := 0; i < n; i++ {
//		if v, ok := m[nums[i]]; ok {
//			if (v + 1) > n/2 {
//				return nums[i]
//			}
//			m[nums[i]]++
//		} else {
//			m[nums[i]] = 1
//		}
//	}
//	return 0
//}

// func2 投票选拔
//func majorityElement(nums []int) int {
//	var major = 0
//	var count = 0
//
//	for _, v := range nums {
//		if count == 0 {
//			major = v
//		}
//		if major == v {
//			count++
//		} else {
//			count--
//		}
//	}
//	return major
//}

// func3 排序
func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var n = len(nums)
	return nums[n/2]
}
