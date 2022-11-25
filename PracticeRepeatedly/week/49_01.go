package main

import "sort"

// 49 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var diff [26]int
		for _, v := range str {
			diff[v-'a']++
		}
		m[diff] = append(m[diff], str)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 排序
func groupAnagrams02(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] > bytes[j]
		})
		m[string(bytes)] = append(m[string(bytes)], str)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
