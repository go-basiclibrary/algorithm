package main

import "sort"

// 88 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
func merge02(nums1 []int, m int, nums2 []int, n int) {
	sorts := make([]int, 0, len(nums1))
	l, r := 0, 0
	for {
		if l == m {
			sorts = append(sorts, nums2[r:]...)
			break
		}
		if r == n {
			sorts = append(sorts, nums1[l:m]...)
			break
		}

		if nums1[l] > nums2[r] {
			sorts = append(sorts, nums2[r])
			r++
		} else {
			sorts = append(sorts, nums1[l])
			l++
		}
	}
	copy(nums1, sorts)
}

// 逆序双指针
func merge03(nums1 []int, m int, nums2 []int, n int) {
	for n1, n2, tail := m-1, n-1, len(nums1)-1; n1 >= 0 || n2 >= 0; tail-- {
		if n1 == -1 {
			nums1[tail] = nums2[n2]
			n2--
		} else if n2 == -1 {
			nums1[tail] = nums1[n1]
			n1--
		} else if nums1[n1] > nums2[n2] {
			nums1[tail] = nums1[n1]
			n1--
		} else {
			nums1[tail] = nums2[n2]
			n2--
		}
	}
}
