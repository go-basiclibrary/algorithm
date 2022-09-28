package main

import "fmt"

// leetcode 01 两数之和
func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], k}
		}
		m[v] = k
	}
	return nil
}
