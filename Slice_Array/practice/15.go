package main

import (
	"fmt"
	"sort"
)

// 15 三数之和
func main() {
	fmt.Println(threeSum02([]int{-1, 0, 1, 2, -1, -4}))
}

/*
	方法一:暴力
	方法二:排序加双指针夹逼
	方法三:引入map(待思考)
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}

func threeSum02(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	if len(nums) < 3 {
		return res
	}
	// 枚举a
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
