package main

// 367 有效的完全平方数
func isPerfectSquare(num int) bool {
	l, r := 0, num/2
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r*r == num {
		return true
	} else {
		return false
	}
}

// 优化
func isPerfectSquare02(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > num {
			r = mid - 1
		} else if mid*mid < num {
			l = mid + 1
		} else { // mid ^ 2 == num
			return true
		}
	}

	return false
}
