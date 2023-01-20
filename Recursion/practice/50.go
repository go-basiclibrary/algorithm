package main

// 50 Pow(x,n)  x的n次方
// 暴力解法
func myPow(x float64, n int) float64 {
	var result float64
	result = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

// 递归解法(分治)
func myPowDFS(x float64, n int) float64 {
	// 考虑负数情况
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
