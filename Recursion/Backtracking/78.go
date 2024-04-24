package main

// 子集

// 回溯
//func subsets(nums []int) [][]int {
//	var ans [][]int
//	var temp = make([]int, 0, len(nums))
//
//	var dfs func(index int)
//	dfs = func(index int) {
//		// terminator
//		if index == len(nums) {
//			ans = append(ans, append(make([]int, 0, len(temp)), temp...))
//			return
//		}
//		// current logic
//		// choose or no choose
//		temp = append(temp, nums[index])
//		// drill down
//		dfs(index + 1)
//		// no choose 回溯
//		temp = temp[:len(temp)-1]
//		dfs(index + 1)
//	}
//
//	// call
//	dfs(0)
//	return ans
//}
