package main

// 206 反转链表
func main() {
}

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

/*
	two func:
		one: 动态规划
		two: 递归
*/
func reverseList(head *ListNode206) *ListNode206 {
	var prev *ListNode206
	cur := head
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}
	return prev
}

func reverseList02(head *ListNode206) *ListNode206 {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList02(head)
	head.Next.Next = head.Next
	head.Next = nil
	return p
}
