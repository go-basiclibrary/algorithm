package practice

//func combine(n int, k int) (Ans [][]int) {
//	temp := []int{}
//	var dfs func(int)
//	dfs = func(cur int) {
//		// 剪枝: temp长度加上区间[cur,n]的长度小于k,不可能构造出长度为k的temp
//		if len(temp)+(n-cur+1) < k {
//			return
//		}
//		// 记录合法答案
//		if len(temp) == k {
//			comb := make([]int, k)
//			copy(comb, temp)
//			Ans = append(Ans, comb)
//			return
//		}
//		// 考虑选择当前位置
//		temp = append(temp, cur)
//		dfs(cur + 1)
//		temp = temp[:len(temp)-1]
//		// 考虑不选择当前位置
//		dfs(cur + 1)
//	}
//	dfs(1)
//	return
//}

func combine(n int, k int) (ans [][]int) {
	temp := make([]int, 0, k)
	var dfs func(cur int)
	dfs = func(cur int) {
		// 如果长度不够了退出
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 如果temp达到了要求长度,也可以进行退出
		if len(temp) == k {
			res := make([]int, k)
			// 基于temp数据会变,这里需要临时变量去加入ans
			copy(res, temp)
			ans = append(ans, res)
			return
		}
		// 考虑当前该数
		temp = append(temp, cur)
		dfs(cur + 1)
		// 不考虑当前数,进行回溯
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}
