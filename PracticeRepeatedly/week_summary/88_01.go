package main

import "sort"

// 88 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
func merge02(nums1 []int, m int, nums2 []int, n int) {
	newList := make([]int, m+n)
	l, r := 0, 0

	for {
		if l == n {
			newList = append(newList, nums2[r:]...)
			break
		}
		if r == m {
			newList = append(newList, nums2[l:]...)
		}

		if nums1[l] < nums2[r] {
			newList = append(newList, nums1[l])
			l++
		} else {
			newList = append(newList, nums2[r])
			r++
		}
	}

	copy(nums1, newList)
}

// 尾插法
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
