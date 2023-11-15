package main

import (
	"fmt"
)

func main() {
	res := isValid("]")
	fmt.Println(res)
}

// 有效括号
func isValid(s string) bool {
	var m = map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	var queue = make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		// 如果是左符号进行入栈
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			queue = append(queue, s[i])
		} else { // 如果是右符号则进行出栈
			// 进行比较是否匹配
			if len(queue) != 0 && m[s[i]] != queue[len(queue)-1] { // 不匹配退出
				return false
			} else if len(queue) != 0 && m[s[i]] == queue[len(queue)-1] { // 匹配则进行出栈
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		}
	}
	// 如果结束后,queue依然存在字符,则该串不符合要求
	if len(queue) != 0 {
		return false
	}
	return true
}

// 优化条件
func isValid02(s string) bool {
	var m = map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	var queue = make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		// 如果是左符号进行入栈
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			queue = append(queue, s[i])
		} else { // 如果是右符号则进行出栈
			// 进行比较是否匹配
			if len(queue) == 0 || m[s[i]] != queue[len(queue)-1] {
				return false
			}
			queue = queue[:len(queue)-1]
		}
	}

	return len(queue) == 0
}
