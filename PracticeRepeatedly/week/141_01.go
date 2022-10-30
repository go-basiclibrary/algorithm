package main

import "fmt"

// 141	环状链表
func main() {
}

// map 快慢 地址
// map
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle02(head *ListNode) bool {
	l, r := head, head
	for l != nil && r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
		if l == r {
			return true
		}
	}
	return false
}

// 地址
func hasCycle03(head *ListNode) bool {
	for head != nil && head.Next != nil {
		if fmt.Sprintf("%p", head) >= fmt.Sprintf("%p", head.Next) {
			return true
		}
		head = head.Next
	}

	return false
}
