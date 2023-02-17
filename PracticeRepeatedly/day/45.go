package main

// 45 跳跃游戏Ⅱ
func jump(nums []int) int {
	var rightStep int
	var step, end int
	for i := 0; i < len(nums)-1; i++ { // 不算最后一个格子,因为有恰好到达最后的情况,此时算会多走一步
		if i <= rightStep { // rightStep大于i表示我当前依然可以走到的i范围内
			rightStep = maxR(rightStep, nums[i]+i)
			if i == end {
				end = rightStep
				step++
			}
		}
	}
	return step
}

func maxR(step int, i int) int {
	if step > i {
		return step
	}
	return i
}
