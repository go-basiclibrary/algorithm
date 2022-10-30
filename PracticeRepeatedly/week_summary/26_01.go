package main

// 26 删除数组中的重复项
// 双指针,控制
func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	j++
	return j
}
