package main

import "fmt"

// 142 环状链表Ⅱ
// 数学 map 地址
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

//数学
func detectCycle02(head *ListNode) *ListNode {
	//第一次相遇,即再走a步两节点再次相遇
	slow, quick := head, head
	for slow != nil && quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			quick = head
			break
		}
	}

	if quick == nil || quick.Next == nil {
		return nil
	}

	for quick != slow {
		quick = quick.Next
		slow = slow.Next
	}
	return quick
}

// 地址
func detectCycle03(head *ListNode) *ListNode {
	for head != nil && head.Next != nil {
		if fmt.Sprintf("%p", head) >= fmt.Sprintf("%p", head.Next) {
			return head.Next
		}
		head = head.Next
	}

	return nil
}
