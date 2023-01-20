package main

import "sort"

// 242 有效的字母异位词
func isAnagram(s string, t string) bool {
	var diff [26]int
	var diff2 [26]int
	for _, v := range s {
		diff[v-'a']++
	}
	for _, v := range t {
		diff2[v-'a']++
	}
	if diff == diff2 {
		return true
	}
	return false
}

// 优化空间
func isAnagram02(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var diff [26]int
	sign := 0
	for _, v := range s {
		diff[v-'a']++
		sign++
	}
	for _, v := range t {
		if diff[v-'a'] > 0 {
			diff[v-'a']--
			sign--
		}
	}
	if sign == 0 {
		return true
	}
	return false
}

// map解法
func isAnagram03(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[int32]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

// sort 解法
func isAnagram04(s string, t string) bool {
	var s1, s2 = []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] > s2[j]
	})
	if string(s1) == string(s2) {
		return true
	}
	return false
}
