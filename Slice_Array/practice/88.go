package main

import (
	"sort"
)

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

// 88 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	//考虑,n 是零的情况
	if n == 0 {
		return
	}

	// 合并
	copy(nums1[m:], nums2[:])
	// 再排序
	sort.Ints(nums1)
}

// 双指针
func merge02(nums1 []int, m int, nums2 []int, n int) {
	sorts := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorts = append(sorts, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorts = append(sorts, nums1[p1:m+1]...)
			break
		}

		if nums1[p1] > nums2[p2] {
			sorts = append(sorts, nums2[p2])
			p2++
		} else {
			sorts = append(sorts, nums1[p1])
			p1++
		}
	}

	copy(nums1, sorts)
}

// 逆序双指针
func merge03(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, len(nums1)-1; p1 >= 0 || p2 >= 0; tail-- {
		if p1 == -1 {
			nums1[tail] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[tail] = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
	}
}
