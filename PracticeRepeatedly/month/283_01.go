package main

// 283 移动零
// 双指针
func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left]
			if nums[left] != 0 {
				left++
			}
		}
	}
}

// 单一变量控制
func moveZeroes02(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[left] = nums[i]
			if i != left { //说明前面遇到了零 i在++ left没++
				nums[i] = 0
			}
			left++
		}
	}
}

// 滚雪球
func moveZeroes03(nums []int) {
	snowSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i-snowSize], nums[i] = nums[i], nums[i-snowSize]
		} else {
			snowSize++ // 记录雪球数
		}
	}
}
