package main

// 1 两数之和
func main() {
}

/*
	方法一:暴力求解
	方法二:引入map,只需一层for循环即可,空间换时间
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

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
