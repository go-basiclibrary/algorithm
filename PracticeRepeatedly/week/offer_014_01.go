package main

//Offer 014 字符中的变为词
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff1 [26]int
	var diff2 [26]int
	for i, v := range s1 {
		diff1[v-'a']++
		diff2[s2[i]-'a']++
	}
	if diff1 == diff2 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		diff2[s2[i-len(s1)]-'a']--
		diff2[s2[i]-'a']++
		if diff1 == diff2 {
			return true
		}
	}
	return false
}

// 维护加前 加后
func checkInclusion02(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff [26]int
	var sign int
	for i, v := range s1 {
		diff[v-'a']++
		diff[s2[i]-'a']--
	}
	for _, v := range diff {
		if v != 0 {
			sign++
		}
	}
	if sign == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		l, r := s2[i-len(s1)]-'a', s2[i]-'a' //出栈,入栈
		if l == r {
			continue
		}
		if diff[l] == 0 {
			sign++
		}
		diff[l]++
		if diff[l] == 0 {
			sign--
		}
		if diff[r] == 0 {
			sign++
		}
		diff[r]--
		if diff[r] == 0 {
			sign--
		}
		if sign == 0 {
			return true
		}
	}
	return false
}
