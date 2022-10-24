package main

// ListNode24 反转链表
type ListNode24 struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}

	return prev
}

// 递归
func reverseList02(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
