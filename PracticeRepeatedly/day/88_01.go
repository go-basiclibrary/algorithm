package main

import "sort"

// 88  合并两个有序数组
//内置方法
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

//双指针
func merge02(nums1 []int, m int, nums2 []int, n int) {
	sorts := make([]int, 0, m+n)
	l, r := 0, 0
	for {
		if l == m {
			sorts = append(sorts, nums2[r:]...)
			break
		}
		if r == n {
			sorts = append(sorts, nums1[l:m+1]...)
			break
		}

		if nums1[l] < nums2[r] {
			sorts = append(sorts, nums1[l])
			l++
		} else {
			sorts = append(sorts, nums2[r])
			r++
		}
	}

	copy(nums1, sorts)
}

// 逆序双指针(尾插法)
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
