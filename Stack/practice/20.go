package main

// 20 有效的括号
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			//入栈
			stack = append(stack, s[i])
		} else {
			//出栈比较
			if len(stack) == 0 || pairs[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
