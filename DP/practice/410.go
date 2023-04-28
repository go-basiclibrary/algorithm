package main

import "math"

func splitArray(nums []int, k int) int {
	n := len(nums)
	f := make([][]int, n+1)
	sub := make([]int, n+1)

	for i := 0; i < len(f); i++ {
		f[i] = make([]int, k+1)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, k); j++ {
			for m := 0; m < i; m++ {
				f[i][j] = min(f[i][j], max(f[m][j-1], sub[i]-sub[m]))
			}
		}
	}
	return f[n][k]
}
