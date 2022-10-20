package main

import "sort"

// 15 三数之和
func main() {
}

/*
	two func
		one:暴力  3for 但被特殊条件过滤  0,0,0
		two:双指针夹逼
*/
func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 || len(nums) < 3 {
			return res
		}
		// 过滤相同元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 需要遍历的头尾
		head := i + 1
		tail := len(nums) - 1
		for head < tail {
			if nums[i]+nums[head]+nums[tail] == 0 {
				// 这里找到要两数同时+-
				res = append(res, []int{nums[i], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] { // 循环查找下面的数,并过滤重复
					head++
				}
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				head++
				tail--
			} else if nums[i]+nums[head]+nums[tail] > 0 {
				tail--
			} else {
				head++
			}
		}
	}

	return res
}
