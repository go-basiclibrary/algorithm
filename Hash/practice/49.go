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

func myHash(s string) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19,
		23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89,
		97, 101} // 用质数表示26个字母

	ans := 1 // 初始化哈希值为1
	for _, b := range s {
		ans *= primes[b-'a'] // 每个字母对应一个质数，相乘得到哈希值
		ans %= 1e9 + 7       // 防止溢出
	}
	return ans
}

// func,使用质数hash计算方式
func groupAnagrams03(strs []string) [][]string {
	mp := map[int][]string{} // 键为哈希值，值为字符串切片
	for _, s := range strs {
		h := myHash(s)           // 计算哈希值
		mp[h] = append(mp[h], s) // 将字符串加入对应的键中
	}
	ans := make([][]string, 0, len(mp)) // 初始化返回值为一个二维字符串切片
	for _, v := range mp {
		ans = append(ans, v) // 将哈希值对应的字符串切片加入返回值
	}
	return ans
}
