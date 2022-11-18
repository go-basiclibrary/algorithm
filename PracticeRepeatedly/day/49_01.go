package main

import "sort"

// 49 字母异位词分组
// 方法一:排序,map
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]string)
	for _, v := range strs {
		v1 := []byte(v)
		sort.Slice(v1, func(i, j int) bool {
			return v1[i] > v1[j]
		})
		m[string(v1)] = append(m[string(v1)], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 比较数组,map
func groupAnagrams02(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[[26]int][]string)
	for _, v := range strs {
		var data [26]int
		for _, i := range v {
			data[i-'a']++
		}
		m[data] = append(m[data], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
