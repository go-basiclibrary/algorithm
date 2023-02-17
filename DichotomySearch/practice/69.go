package main

// 69 x的平方根
// 二分
func mySqrt(x int) int {
	// terminate
	if x == 0 || x == 1 {
		return x
	}
	// dichotomy search
	l, r := 0, x/2
	for l <= r {
		mid := (r + l) / 2
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

// 牛顿迭代
func mySqrtND(x int) int {
	r := x / 2
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
