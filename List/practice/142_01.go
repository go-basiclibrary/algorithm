package main

import "fmt"

// 142 环形链表Ⅱ
func main() {
}

type ListNode142 struct {
	Val  int
	Next *ListNode142
}

/*
	三种解法
		map
		双指针
		链表底层地址从低到高
*/
// 链表底层地址从低到高分配
func detectCycle01(head *ListNode142) *ListNode142 {
	for head != nil {
		if fmt.Sprintf("%p", head) >= fmt.Sprintf("%p", head.Next) {
			return head.Next
		}
		head = head.Next
	}

	return nil
}

// map解法
func detectCycle02(head *ListNode142) *ListNode142 {
	m := make(map[*ListNode142]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return nil
}

// 数学算法
func detectCycle03(head *ListNode142) *ListNode142 {
	quick, slow := head, head
	for quick != nil && slow != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			quick = head //去头
			break
		}
	}
	if quick == nil || quick.Next == nil {
		return nil
	}
	for quick != slow {
		quick, slow = quick.Next, slow.Next
	}
	return quick
}
