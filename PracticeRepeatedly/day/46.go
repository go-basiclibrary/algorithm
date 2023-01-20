package main

var ans [][]int

// 46 全排列
func permute(nums []int) [][]int {
	ans = [][]int{} // 每回合初始化
	dfs(nums, len(nums), make([]int, 0, len(nums)))
	return ans
}

func dfs(nums []int, k int, path []int) {
	if len(nums) == 0 {
		// path写满,可以写入一个答案
		res := make([]int, len(path))
		copy(res, path)
		ans = append(ans, res)
	}
	for i := 0; i < k; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) // i数据从nums中去除
		dfs(nums, len(nums), path)
		//回归数据
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}
