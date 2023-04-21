package main

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	ret := math.MinInt
	for i := 0; i < m; i++ {
		mergeRow := make([]int, n)
		for j := i; j < m; j++ {
			for col := 0; col < n; col++ {
				mergeRow[col] += matrix[j][col]
			}
			// 不超过K的最大连续子数组
			for x := 0; x < n; x++ {
				var tempS int
				for y := x; y < n; y++ {
					tempS += mergeRow[y]
					if tempS == k {
						return k
					} else if tempS < k {
						ret = max(ret, tempS)
					}
				}
			}
		}
	}

	return ret
}

func max(ret int, s int) int {
	if ret > s {
		return ret
	}
	return s
}
