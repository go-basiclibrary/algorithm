package main

import "sort"

// 242 有效的字母异位词
// 双数组解决
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var diff1 [26]int
	var diff2 [26]int
	for _, v := range s {
		diff1[v-'a']++
	}
	for _, v := range t {
		diff2[v-'a']++
	}
	if diff1 == diff2 {
		return true
	}
	return false
}

// 优化
func isAnagram02(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var diff [26]int
	for _, v := range s {
		diff[v-'a']++
	}
	for _, v := range t {
		diff[v-'a']--
	}
	for _, v := range diff {
		if v != 0 {
			return false
		}
	}
	return true
}

// map
func isAnagram03(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
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

// sort
func isAnagram04(s string, t string) bool {
	s1 := []byte(s)
	t1 := []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] > t1[j]
	})

	return string(s1) == string(t1)
}
