package main

//283  移动零
// 1 滚雪球
func moveZeroes(nums []int) {
	zeroSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroSize++
		} else if zeroSize > 0 {
			nums[i], nums[i-zeroSize] = nums[i-zeroSize], nums[i]
		}
	}
}

// 双指针
func moveZeroes02(nums []int) {
	l, r := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

// 单指针
func moveZeroes03(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}
