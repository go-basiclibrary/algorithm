package main

// offer_014 剑指 Offer II 014. 字符串中的变位词
// 方法1,使用切片,查看内部数量是否为0(滑动法)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff [26]int // s1存在++ s2存在++
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
		if diff == diff2 { // 需要优化的点,因为每次都进行了比较
			return true
		}
	}

	return false
}

// 方法2 对于方法一多次比较进行优化
func checkInclusion02(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var btn [26]int // s1存在-- s2存在++
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

// 方法3 对于方法二多次比较diff进行优化
func checkInclusion03(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var diff [26]int
	for i := 0; i < len(s1); i++ {
		diff[s1[i]-'a']-- //保持diff一直为负数
	}
	left := 0
	for right, v := range s2 {
		x := v - 'a'
		// diff中,存在>0那么right必须右移
		diff[x]++
		for diff[x] > 0 {
			diff[s2[left]-'a']--
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
	}
	return false
}
