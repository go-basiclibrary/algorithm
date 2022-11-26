package main

// offer 014
// 双指针解法
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff [26]int
	for _, v := range s1 {
		diff[v-'a']--
	}
	l := 0 // 左指针
	for r, v := range s2 {
		diff[v-'a']++
		// 如果diff中存在大于零的,那么left要出去
		for diff[v-'a'] > 0 { //大于零就要移动出去
			diff[s2[l]-'a']--
			l++
		}
		//如果存在diff都等于零,且长度相等,主要依赖于长度相等
		if r-l+1 == len(s1) {
			return true
		}
	}
	return false
}
