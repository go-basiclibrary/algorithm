package main

// 55 跳跃游戏
func canJump(nums []int) bool {
	var rightmost int
	for i := 0; i < len(nums); i++ {
		if i <= rightmost {
			rightmost = maxRight(rightmost, i+nums[i])
			if rightmost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

func maxRight(rightmost int, i int) int {
	if rightmost > i {
		return rightmost
	}
	return i
}
