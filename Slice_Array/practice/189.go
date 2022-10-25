package main

func main() {
	rotate03([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// 189 轮转数组
// 1.开新数组
func rotate(nums []int, k int) {
	newArray := make([]int, len(nums))
	for i, v := range nums {
		newArray[(i+k)%len(nums)] = v
	}
	copy(nums, newArray)
}

// 2.反转数组
func rotate02(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

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
