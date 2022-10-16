package main

import (
	"fmt"
	"strings"
)

const testStr = `A man, a plan, a canal: Panama`

// 125 验证字符串是否回文
func main() {
	fmt.Printf("the 125 title answer is %v\n", isPalindrome(testStr))
}

func isPalindrome(s string) bool {
	//大小写转换
	s = strings.ToLower(s)
	//过滤special字符
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			sb.WriteByte(s[i])
		}
	}
	s = sb.String()
	//前后比较
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
