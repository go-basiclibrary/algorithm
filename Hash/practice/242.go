package main

import "sort"

// 242 有效的字母异位词
// 切片求法
func isAnagram(s string, t string) bool {
	var diff [26]int
	var diff2 [26]int
	for i := 0; i < len(s); i++ {
		diff[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		diff2[t[i]-'a']++
	}
	if diff == diff2 {
		return true
	}
	return false
}

// 优化  节省空间
func isAnagram02(s string, t string) bool {
	var diff [26]int
	sign := 0
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		diff[s[i]-'a']++
		sign++
	}

	for i := 0; i < len(t); i++ {
		if diff[t[i]-'a'] > 0 {
			diff[t[i]-'a']--
			sign--
		}
	}
	if sign == 0 {
		return true
	}
	return false
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

//sort
func isAnagram04(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
