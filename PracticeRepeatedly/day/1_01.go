package main

// 11两数之和
func main() {
}

/*
	two func
		one:暴力解法
		two:接入哈希
*/
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

// hash
func twoSum02(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{k, m[target-v]}
		}
		m[v] = k
	}
	return nil
}

func twoSum03(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{i, m[target-nums[i]]}
		}
		m[nums[i]] = i
	}
	return nil
}
