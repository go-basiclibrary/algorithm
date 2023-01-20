package practice

var ans [][]int

// Permute 46 全排列  回溯算法
func Permute(nums []int) [][]int {
	ans = [][]int{} //数据初始
	dfs(nums, len(nums), make([]int, 0, len(nums)))
	return ans
}

func dfs(nums []int, k int, path []int) {
	if len(nums) == 0 {
		res := make([]int, len(path))
		copy(res, path)
		ans = append(ans, res)
	}
	for i := 0; i < k; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //将数据i切出去
		dfs(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯数据回去
		path = path[:len(path)-1]
	}
}
