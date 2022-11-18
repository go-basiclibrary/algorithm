package main

// offer 014
// 双指针
func checkInclusion04(s1 string, s2 string) bool {
	var diff [26]int
	//将s1收集到diff中
	for _, v := range s1 {
		diff[v-'a']-- // diff中一直保持为负
	}
	left := 0
	for right, v := range s2 {
		diff[v-'a']++
		for diff[v-'a'] > 0 {
			diff[s2[left]-'a']-- // 左边出去,要保持diff <=0
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
	}
	return false
}
