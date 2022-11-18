package main

import "sort"

// 49 字母异位词分组
// map解法
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	//排序
	//针对strs中的每一个内容进行排序,有存在的,则存入res
	for _, v := range strs {
		v1 := []byte(v)
		sort.Slice(v1, func(i, j int) bool {
			return v1[i] > v1[j]
		})
		m[string(v1)] = append(m[string(v1)], v)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

//统计字符解法
func groupAnagrams02(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, v := range strs {
		var str [26]int
		for _, i := range v {
			str[i-'a']++
		}
		m[str] = append(m[str], v)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
