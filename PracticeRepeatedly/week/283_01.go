package main

// 283 移动零
// 双指针方式
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if j < i {
				nums[i] = 0
			}
			j++
		}
	}
}

// 左右指针
func moveZeroes02(nums []int) {
	l, r := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			r++
		}
		l++
	}
}

// 滚雪球
func moveZeroes03(nums []int) {
	snowSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowSize++
		} else if snowSize > 0 {
			nums[i], nums[i-snowSize] = nums[i-snowSize], nums[i]
		}
	}
}
