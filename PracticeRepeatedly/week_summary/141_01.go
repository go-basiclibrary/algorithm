package main

// ListNode141 141 环形链表
type ListNode141 struct {
	Val  int
	Next *ListNode
}

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
	quick, slow := head, head
	for quick != nil && slow != nil && slow.Next != nil {
		quick = quick.Next
		slow = slow.Next.Next
		if quick == slow {
			return true
		}
	}
	return false
}
