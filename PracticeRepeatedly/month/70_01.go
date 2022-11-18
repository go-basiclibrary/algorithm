package main

func main() {
	climbStairs(4)
}

// 递归  归纳法,可以知道,该解法就是求菲薄
// 1 1
// 2 2
// 3 3
// 4 5
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
