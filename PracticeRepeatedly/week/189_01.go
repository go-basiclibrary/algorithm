package main

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for i := start; i <= end; i++ {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate02(nums []int, k int) {
	newSlice := make([]int, len(nums))
	for i, v := range nums {
		newSlice[(i+k)%len(nums)] = v
	}

	copy(nums, newSlice)
}

// 环状替代
func rotate03(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
