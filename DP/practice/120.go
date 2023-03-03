package main

import "fmt"

func main() {
	triangle := [][]int{
		{1},
		{1, 2},
	}
	res := minimumTotal(triangle)
	fmt.Println(res)
	fmt.Println(triangle)
}

// 120 三角形最小路径和
// 暴力  工程代码不建议这么写,triangle会在内部被修改
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
