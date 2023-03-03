package main

import "fmt"

func main() {
	res := uniquePaths03(3, 7)
	fmt.Println(res)
}

// 62 不同路径	广度优先算法  (m,n)  超时
func uniquePaths(m int, n int) int {
	var num int
	var x, y = 1, 1
	var route func(x, y int)
	route = func(x, y int) {
		// terminator
		if x == m && y == n {
			num++
			return
		}
		// 向下走
		if x < m {
			route(x+1, y)
		}
		// 向右走
		if y < n {
			route(x, y+1)
		}
	}

	route(x, y)
	return num
}

// DP
func uniquePaths02(m int, n int) int {
	var res = make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		res[i][0] = 1
	}
	for i := 0; i < n; i++ {
		res[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

// DP 倒叙
func uniquePaths03(m int, n int) int {
	var res = make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		res[i][n-1] = 1
	}
	for i := 0; i < n; i++ {
		res[m-1][i] = 1
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			res[i][j] = res[i+1][j] + res[i][j+1]
		}
	}
	return res[0][0]
}
