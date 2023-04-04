package main

import (
	"math"
)

func numSquares(n int) int {
	f := make([]int, n+1) // 表示最少需要多少个数的平方来表示整数i
	for i := 1; i <= n; i++ {
		minn := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}

	return f[n]
}
