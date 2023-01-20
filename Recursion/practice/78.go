package main

func main() {
	subsets([]int{1, 2, 3})
}

// 78 子集
// 回溯递归
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	set := []int{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, set...)) //复制一份值
			return
		}
		dfs(index + 1)
		set = append(set, nums[index])
		dfs(index + 1)
		set = set[:len(set)-1]
	}
	dfs(0)
	return ans
}
