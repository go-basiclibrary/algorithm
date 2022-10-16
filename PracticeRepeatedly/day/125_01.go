package main

import "strings"

//验证字符串是否回文
func main() {

}

// time O(n)  space O(1)
func isPalindrome02(s string) bool {
	//双指针求法  降低空间复杂度
	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// time O(n)  space O(n)  this I think space is O(1)
func isPalindrome(s string) bool {
	//大小写
	s = strings.ToLower(s)
	//过滤特殊字符串  sb解决了空间复杂度的问题
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			sb.WriteByte(s[i])
		}
	}
	s = sb.String()
	//头尾比较
	for i, j := 0, len(s); i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
