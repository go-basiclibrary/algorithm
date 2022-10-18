package main

// 70 爬楼梯
func main() {

}

// 方法:fibo解决
func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
