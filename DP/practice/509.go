package main

import "fmt"

// 509 斐波那契数
// 递归
// Time complexity O(n^2) 指数级增长
func fib(n int) int {
	// terminator
	if n <= 1 {
		return n
	}
	// process current logic
	return fib(n-1) + fib(n-2) // 如何避免重复计算
}

var common = make([]int, 50)

// 将已经走过的节点存储下来,简化成O(n),但是我们的内存消耗增加
func fibTwo(n int) int {
	// terminator
	if n <= 1 {
		return n
	}
	// process current logic
	if common[n] == 0 {
		common[n] = fib(n-1) + fib(n-2)
	}
	return common[n]
}

// DP动规 O(n)
func fibThree(n int) int {
	common[0], common[1] = 0, 1
	for i := 2; i <= n; i++ {
		common[i] = common[i-1] + common[i-2]
	}
	return common[n]
}

func main() {
	b := fibTwo(4)
	w := fibTwo(5)
	fmt.Println(b)
	fmt.Println(w)
}
