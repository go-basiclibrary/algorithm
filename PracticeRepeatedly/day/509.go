package main

// 509 斐波那契数列三种解法
// 递归
func fib03(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

var common = make([]int, 32)

// LRU
func fib04(n int) int {
	if n <= 1 {
		return n
	}
	if common[n] == 0 {
		return fib(n-1) + fib(n-2)
	}
	return common[n]
}

// DP 动规
func fib05(n int) int {
	if n <= 1 {
		return n
	}
	common[0], common[1] = 0, 1
	for i := 2; i <= n; i++ {
		if common[i] == 0 {
			common[i] = common[i-1] + common[i-2]
		}
	}
	return common[n]
}
