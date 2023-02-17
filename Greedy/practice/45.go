package main

// 45 跳跃游戏 Ⅱ
func jump(nums []int) int {
	var rightJump int
	var step, end int
	for i := 0; i < len(nums)-1; i++ {
		// 我们贪心的认为,只有最大的步数且可达的步数,才会计算为我们走的一步
		if rightJump >= i {
			rightJump = maxJump(rightJump, nums[i]+i)
			if i == end {
				end = rightJump
				step++
			}
		}
	}
	return step
}

func maxJump(rightJump int, i int) int {
	if rightJump > i {
		return rightJump
	}
	return i
}
