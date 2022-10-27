package main

//1 两数之和
//暴力
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// map
func twoSum02(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if res, ok := m[target-nums[i]]; ok {
			return []int{res, i}
		}
		m[nums[i]] = i
	}
	return nil
}
