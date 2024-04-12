package main

// 70 爬楼梯
//func main() {
//	res := climbStairs(45)
//	fmt.Println(res)
//}

// func1 递归 查找最小重复子问题
// time  超时
// space
//func climbStairs(n int) int {
//	if n <= 3 {
//		return n
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//}

// func2 fibo
// time 0ms 100
// space 1.99MB 25.84
//func climbStairs(n int) int {
//	if n <= 3 {
//		return n
//	}
//	var x, y = 2, 3
//	for i := 4; i <= n; i++ {
//		x, y = y, x+y
//	}
//	return y
//}

// func3 递归 查找最小重复子问题 对重复计算进行保存
// time  0ms 100
// space 2.05MB 13.25
//func climbStairs(n int) int {
//	if n <= 3 {
//		return n
//	}
//	// 进行相同计算值的记录
//	var memo = make(map[int]int, n+1)
//	return climbStairsMemo(n, memo)
//}
//
//func climbStairsMemo(n int, memo map[int]int) int {
//	if v, ok := memo[n]; ok {
//		return v
//	}
//
//	if n == 1 {
//		return n
//	}
//	if n == 2 {
//		return n
//	}
//
//	memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
//
//	return memo[n]
//}
