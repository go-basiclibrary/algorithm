package main

import "sort"

// 49 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		b := []byte(v)
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})
		m[string(b)] = append(m[string(b)], v)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 统计字符
func groupAnagrams02(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, v := range strs {
		var diff [26]int
		for _, i := range v {
			diff[i-'a']++
		}
		m[diff] = append(m[diff], v)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
