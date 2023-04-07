package main

// 55 跳跃游戏
func canJump(nums []int) bool {
	var maxJump int
	for i := 0; i < len(nums); i++ {
		// DP方程
		// 最大跳跃长度 = max(当前位置加当前最大跳跃,上某个位置最大跳跃)
		if i <= maxJump {
			maxJump = max(i+nums[i], maxJump)
			if maxJump+1 >= len(nums) {
				return true
			}
		}
	}
	return false
}
