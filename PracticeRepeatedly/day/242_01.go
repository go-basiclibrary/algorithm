package main

import "sort"

// 242 有效的字母异位词
// array
func isAnagram(s string, t string) bool {
	var diff [26]int
	var diff2 [26]int // 多用空间,要减少空间复杂度
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

// 方法二
func isAnagram02(s string, t string) bool {
	// 降低空间复杂度
	if len(s) != len(t) {
		return true
	}
	var diff [26]int
	var sum int
	for _, v := range s {
		diff[v-'a']++
		sum++
	}
	for _, v := range t {
		if diff[v-'a'] > 0 {
			diff[v-'a']--
			sum--
		}
	}
	if sum == 0 {
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

// sort
func isAnagram04(s string, t string) bool {
	s1, t1 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] > t1[j]
	})
	return string(s1) == string(t1)
}
