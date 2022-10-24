package main

// 70爬楼梯
// fib
//func climbStairs(n int) int {
//	if n < 2 {
//		return n
//	}
//	a, b := 0, 1
//	for i := 1; i <= n; i++ {
//		a, b = b, a+b
//	}
//	return b
//}

// 递归
func climbStairs(n int) int {
	if (n - 1) < 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
