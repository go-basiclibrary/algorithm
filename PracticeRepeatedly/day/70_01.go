package main

// 70 经典爬楼梯
func main() {
}

// 方法:当没有解题思路的时候,我们可以进行常规情况分析法,就是列举所有的情况进行分析  or (暴力)
// 这题肯定无暴力解法
// 1:1
// 2:2
// 3:f(3) = f(2) + f(1)      只有从1走2才能到3 从2走1才能到4,如果想不出来,则继续写常规想法
// example 3: 1,1,1 1,2 2,1  3种情况,就能猜测到fibo
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
