package main

// offer Ⅱ 014 剑指 Offer II 014. 字符串中的变位词
//func checkInclusion(s1 string, s2 string) bool {
//	if len(s1) > len(s2) {
//		return false
//	}
//	var snt1 [26]int
//	var snt2 [26]int
//	for i := 0; i < len(s1); i++ {
//		snt1[s1[i]-'a']++
//		snt2[s2[i]-'a']++
//	}
//	if snt1 == snt2 {
//		return true
//	}
//	for i := len(s1); i < len(s2); i++ {
//		//滑动,一进一出
//		snt2[s2[i]-'a']++
//		snt2[s2[i-len(s1)]-'a']--
//		if snt1 == snt2 {
//			return true
//		}
//	}
//
//	return false
//}

// 优化,较少比较次数
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	btn := [26]int{}
	for i := 0; i < len(s1); i++ {
		btn[s1[i]-'a']--
		btn[s2[i]-'a']++
	}
	diff := 0
	for _, v := range btn {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		x, y := s2[i]-'a', s2[i-len(s1)]-'a'
		if x == y {
			continue
		}
		if btn[x] == 0 {
			diff++
		}
		btn[x]++
		if btn[x] == 0 {
			diff--
		}
		if btn[y] == 0 {
			diff++
		}
		btn[y]--
		if btn[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
