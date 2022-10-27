package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 206 反转链表
// 动态规划
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 递归
func reverseList02(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList02(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
