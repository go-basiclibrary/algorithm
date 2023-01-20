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

// 滑动slice
func checkInclusion02(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff [26]int
	var diff2 [26]int

	for i := 0; i < len(s1); i++ {
		diff[s1[i]-'a']++
		diff2[s2[i]-'a']++
	}

	if diff == diff2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		diff2[s2[i-len(s1)]-'a']--
		diff2[s2[i]-'a']++
		if diff == diff2 {
			return true
		}
	}

	return false
}

// 优化
func checkInclusion03(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff [26]int
	for i := 0; i < len(s1); i++ {
		diff[s1[i]-'a']--
		diff[s2[i]-'a']++
	}
	sign := 0
	for _, v := range diff {
		if v != 0 {
			sign++
		}
	}
	if sign == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		x, y := s2[i]-'a', s2[i-len(s1)]-'a'
		if x == y {
			continue
		}
		if diff[x] == 0 {
			sign++
		}
		diff[x]++
		if diff[x] == 0 {
			sign--
		}
		if diff[y] == 0 {
			sign++
		}
		diff[y]--
		if diff[y] == 0 {
			sign--
		}
		if sign == 0 {
			return true
		}
	}
	return false
}

// 双指针
func checkInclusion04(s1 string, s2 string) bool {
	var diff [26]int
	for _, v := range s1 {
		diff[v-'a']-- // 将s1存在的置为-
	}
	left := 0 // 左标,用于计算后续长度
	for right, v := range s2 {
		diff[v-'a']++
		for diff[v-'a'] > 0 {
			diff[s2[left]-'a']--
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
	}
	return false
}
