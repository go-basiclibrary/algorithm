package main

// 53 最大子数组和
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] { // 如果当前元素,加上前一个元素和大于当前元素,那么当前元素进行变更为最大值
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
