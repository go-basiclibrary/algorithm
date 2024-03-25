package main

// 1 两数之和
func main() {
}

// 方法1,使用hash
// 最佳解法
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		} else {
			m[v] = i
		}
	}
	return nil
}
