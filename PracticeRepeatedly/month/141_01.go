package main

import "fmt"

// 141 环形链表  map,快慢
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

// 数学
func hasCycle02(head *ListNode) bool {
	for head != nil && head.Next != nil {
		if fmt.Sprintf("%p", head) >= fmt.Sprintf("%p", head.Next) {
			return true
		}
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle03(head *ListNode) bool {
	s, k := head, head
	for s != nil && k != nil && k.Next != nil {
		s = s.Next
		k = k.Next.Next
		if s == k {
			return true
		}
	}
	return false
}
