package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	rotate04(ints, 2)
	fmt.Println(ints)
}

// 189 轮转数组  环状替代
func rotate04(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd1(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd1(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
