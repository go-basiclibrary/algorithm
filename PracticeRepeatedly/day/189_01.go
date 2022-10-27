package main

// 189 轮转数组
// 引入新切片,copy回去,性能低下
func rotate(nums []int, k int) {
	newArray := make([]int, len(nums))
	for i, v := range nums {
		newArray[(k+i)%len(nums)] = v
	}
	copy(nums, newArray)
}

// 通过观察可以多次反转数组
func rotate02(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

// 反转
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
