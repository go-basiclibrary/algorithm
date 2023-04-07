package main

func jump(nums []int) int {
	var maxJump int
	var step, end int
	for i := 0; i < len(nums)-1; i++ { // 不计算最后一次位置,否则会多算一步
		// DP 最大跳跃 = max(最大跳跃,当前位置+最大跳跃)
		maxJump = max(maxJump, i+nums[i])
		// 判断如果i的位置和end位置相等,我们计算出了第一次跳跃后再次最大跳跃新的地点,可以更新该值,并step++
		if i == end { // 当i < end的时候,我们去计算该范围内新的最大跳跃地点
			end = maxJump // 更新end为下一次的maxJump
			step++
		}
	}
	return step
}
