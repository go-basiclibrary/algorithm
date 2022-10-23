package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 地址求解
func detectCycle(head *ListNode) *ListNode {
	for head != nil {
		if fmt.Sprintf("%p", head) >= fmt.Sprintf("%p", head.Next) {
			return head.Next
		}
		head = head.Next
	}

	return nil
}

// map
func detectCycle02(head *ListNode) *ListNode {
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

// 数学求解
func detectCycle03(head *ListNode) *ListNode {
	quick, slow := head, head
	for quick != nil && slow != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
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

	return slow
}
