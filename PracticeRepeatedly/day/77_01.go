package main

// 77 组合
func combine(n int, k int) (ans [][]int) {
	temp := make([]int, 0, k)
	var dfs func(cur int)
	dfs = func(cur int) {
		// 退出条件
		if len(temp)+(n-cur+1) < k {
			return
		}
		if len(temp) == k {
			// 收集结果
			res := make([]int, k)
			copy(res, temp)
			ans = append(ans, res)
			return
		}
		// 考虑当前该数
		temp = append(temp, cur)
		dfs(cur + 1)
		// 不考虑当前数
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}
