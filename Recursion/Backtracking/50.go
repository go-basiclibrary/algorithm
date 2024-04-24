package main

// 50 Pow(x,n)

// func1 回溯,拆解最小子问题
// Olog(n)
//func myPow(x float64, n int) float64 {
//	if n >= 0 {
//		return quickMul(x, n)
//	}
//	return 1.0 / quickMul(x, -n)
//}
//
//func quickMul(x float64, n int) float64 {
//	// terminator
//	if n == 0 {
//		return 1
//	}
//
//	half := myPow(x, n/2)
//
//	if n%2 == 1 {
//		// 奇数
//		return half * half * x
//	} else {
//		return half * half
//	}
//}

// func2 取消递归栈,直接迭代
//func myPow(x float64, n int) float64 {
//	if n >= 0 {
//		return quickMul(x, n)
//	}
//	return 1.0 / quickMul(x, -n)
//}
//
//func quickMul(x float64, n int) float64 {
//	ans := 1.0
//	contribute := x
//	for n > 0 {
//		if n%2 == 1 {
//			ans *= contribute
//		}
//		contribute *= contribute
//		n /= 2
//	}
//	return ans
//}
