package main

import "fmt"

func main() {
	res := climbStairs(3)
	fmt.Println(res)
}

// 爬楼梯 DP
func climbStairs(n int) int {
	var res = make([]int, n)
	if n <= 2 {
		return n
	}
	res[0], res[1] = 1, 2
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[len(res)-1]
}

// DP
func climbStairs02(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
