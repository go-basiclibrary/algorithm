package main

// 70 爬楼梯
func climbStairs(n int) int {
	//fibo
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
