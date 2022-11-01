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
