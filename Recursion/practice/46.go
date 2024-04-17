package main

// 全排列

func main() {
	//path := make([]int, 1)
	//onPath := make([]bool, 2)
	//fmt.Println(path, onPath)
}

// func1 递归 回溯
//func permute(nums []int) [][]int {
//	n := len(nums)
//	path := make([]int, n)
//	onPath := make([]bool, n)
//
//	var ans = make([][]int, 0, fibo(n))
//	var dfs func(int)
//	dfs = func(i int) {
//		if i == n {
//			ans = append(ans, append(make([]int, 0, n), path...))
//			return
//		}
//		for j, on := range onPath {
//			if !on {
//				// chose current j
//				path[i] = nums[j]
//				onPath[j] = true
//				dfs(i + 1)
//				// no 回溯 & no chose j
//				onPath[j] = false
//			}
//		}
//	}
//	return ans
//}
//
//func fibo(n int) int {
//	if n <= 3 {
//		return n
//	}
//	var x, y = 2, 3
//	for i := 4; i <= n; i++ {
//		x, y = y, x+y
//	}
//	return y
//}
//
//var ans [][]int
//
//// func2 另一种回溯
//func permute(nums []int) [][]int {
//	n := len(nums)
//	ans = make([][]int, 0, fibo(n))
//	dfs(nums, n, make([]int, 0, n))
//	return ans
//}
//
//func dfs(dNums []int, k int, path []int) {
//	if len(dNums) == 0 {
//		ans = append(ans, append(make([]int, 0, len(path)), path...))
//	}
//	for i := 0; i < k; i++ {
//		// choose cur or no choose cur
//		cur := dNums[i]
//		// choose cur
//		path = append(path, cur)
//		// clear cur
//		dNums = append(dNums[:i], dNums[i+1:]...)
//		dfs(dNums, len(dNums), path)
//		// no choose cur
//		// 回溯
//		dNums = append(dNums[:i], append([]int{cur}, dNums[i:]...)...)
//		path = path[:len(path)-1]
//	}
//}
