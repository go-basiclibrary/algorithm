package main

// 242 有效的字母异位词
// 切片求法
// 242 有效的字母异位词

func main() {
	//// 测试1
	//result := isAnagram("anagram", "nagaram")
	//fmt.Println(result)
	//// 测试2
	//result = isAnagram("rat", "car")
	//fmt.Println(result)
}

// 方法一,对字符进行排序,并进行比较
//func isAnagram(s string, t string) bool {
//	var sSlice = make([]byte, len(s))
//	var tSlice = make([]byte, len(t))
//
//	for i := 0; i < len(s); i++ {
//		sSlice[i] = s[i]
//	}
//	for i := 0; i < len(t); i++ {
//		tSlice[i] = t[i]
//	}
//	sort.Slice(sSlice, func(i, j int) bool {
//		return sSlice[i] > sSlice[j]
//	})
//	sort.Slice(tSlice, func(i, j int) bool {
//		return tSlice[i] > tSlice[j]
//	})
//	var s1 strings.Builder
//	var t1 strings.Builder
//	s1.Write(sSlice)
//	t1.Write(tSlice)
//
//	return s1.String() == t1.String()
//}

// 方法二,使用map进行一次性比较
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	var m = make(map[byte]int)
//	for i := 0; i < len(s); i++ {
//		if _, ok := m[s[i]]; ok {
//			m[s[i]]++
//		} else {
//			m[s[i]] = 1
//		}
//
//		if _, ok := m[t[i]]; ok {
//			m[t[i]]--
//		} else {
//			m[t[i]] = -1
//		}
//	}
//
//	for _, v := range m {
//		if v != 0 {
//			return false
//		}
//	}
//	return true
//}

// 方法三,优化方法1
// string->byteSlice->string会有copy消耗
//func isAnagram(s string, t string) bool {
//	var s1, t1 = []byte(s), []byte(t)
//	sort.Slice(s1, func(i, j int) bool {
//		return s1[i] > s1[j]
//	})
//	sort.Slice(t1, func(i, j int) bool {
//		return t1[i] > t1[j]
//	})
//	return string(s1) == string(t1)
//}

// 方法四,使用非指针类型hash 最优解
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	var s1 [26]int
//	for i := 0; i < len(s); i++ {
//		s1[s[i]-'a']++
//		s1[t[i]-'a']--
//	}
//	return s1 == [26]int{}
//}

// 方法五,使用指针类型hash
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	var m = make(map[rune]int)
//	for _, v := range s {
//		m[v]++
//	}
//	for _, v := range t {
//		m[v]--
//		if m[v] < 0 {
//			return false
//		}
//	}
//	return true
//}
