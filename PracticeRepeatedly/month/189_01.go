package main

// 189 轮转数组
func rotate(nums []int, k int) {
	res := make([]int, 0, len(nums))
	for i, v := range nums {
		res[(k+i)%len(nums)] = v
	}
	copy(nums, res)
}

// 翻转数组 整体,局部
func rotate02(nums []int, k int) {
	k %= len(nums)
	reverseSlice(nums, 0, len(nums)-1)
	reverseSlice(nums, 0, k-1)
	reverseSlice(nums, k, len(nums)-1)
}

func reverseSlice(nums []int, sta int, end int) {
	for sta < end {
		nums[sta], nums[end] = nums[end], nums[sta]
		sta++
		end--
	}
}
