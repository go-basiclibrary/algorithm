package main

// 1 两数之和
// 暴力解法 O(n^2)
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

// 优化map
func twoSum02(nums []int, target int) []int {
	res := make([]int, 0, 2)
	m := make(map[int]int)
	for k, v := range nums {
		if t, ok := m[target-v]; ok {
			res = append(res, k, t)
			return res
		}
		m[v] = k
	}
	return res
}
