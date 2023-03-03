package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 1, 3}, 3))
}

// 33 搜索旋转排序数组
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] { // 搜索区间
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] { // 搜索区间
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
