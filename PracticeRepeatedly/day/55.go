package main

// 55 跳跃游戏
func canJump(nums []int) bool {
	var now int
	// 覆盖走法,动态规划
	for i := 0; i < len(nums); i++ {
		if i <= now {
			now = maxJump(now, i+nums[i])
			if now >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

func maxJump(now int, i int) int {
	if now > i {
		return now
	}
	return i
}
