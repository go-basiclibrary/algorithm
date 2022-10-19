package main

// 283  移动零
func main() {
	moveZeroes03([]int{0, 1, 0, 3, 12})
}

// 单一变量控制0
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]

			if i != j { // 说明前面遇到零了,exchange
				nums[i] = 0
			}

			j++
		}
	}
}

// 滚雪球
func moveZeroes02(nums []int) {
	snowSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowSize++
		} else if snowSize > 0 {
			nums[i], nums[i-snowSize] = nums[i-snowSize], nums[i]
		}
	}
}

// 双指针解法
func moveZeroes03(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++ //left被交换
		}
		right++
	}
}
