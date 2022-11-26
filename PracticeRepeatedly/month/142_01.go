package main

import "fmt"

// 142 环形链表 Ⅱ
// map
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 地址
func detectCycle02(head *ListNode) *ListNode {
	for head != nil {
		if fmt.Sprintf("%p", head) >= fmt.Sprintf("%p", head.Next) {
			return head.Next
		}
		head = head.Next
	}
	return nil
}

// 双指针
func detectCycle03(head *ListNode) *ListNode {
	s, q := head, head
	for s != nil && q != nil && q.Next != nil {
		if s == q {
			q = head
			break
		}
		s = s.Next
		q = q.Next.Next
	}
	if q == nil || q.Next == nil {
		return nil
	}
	//单步走
	for s != q {
		s = s.Next
		q = q.Next
	}
	return s
}
