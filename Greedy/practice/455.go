package main

import (
	"sort"
)

func main() {
	findContentChildren([]int{3, 2, 1}, []int{3, 2, 1})
}

// 455 分发饼干
func findContentChildren(g []int, s []int) int {
	var res int
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && g[i] > s[j] { // 当前状况满足不了小孩
			j++
		}
		if j < n {
			res++
			j++
		}
	}
	return res
}
