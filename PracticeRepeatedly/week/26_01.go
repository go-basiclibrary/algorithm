package main

func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	j++
	return j
}
