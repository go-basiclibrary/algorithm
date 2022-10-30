package main

// 189 轮转数组
// 三次reverse 新数组 数学
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for i := start; i < end; i++ {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate02(nums []int, k int) {
	newSlice := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newSlice[(k+i)%len(nums)] = nums[i]
	}
	copy(nums, newSlice)
}
