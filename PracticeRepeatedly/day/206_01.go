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

//func reverseList02(head *ListNode206) *ListNode206 {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	p := reverseList02(head)
//	head.Next.Next = head.Next
//	head.Next = nil
//	return p
//}

func reverseList02(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList02(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

// 妖魔化双指针
func reverseList03(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for head.Next != nil {
		t := head.Next.Next //保存下一个起始节点
		head.Next.Next = cur
		cur = head.Next
		head.Next = t
	}
	return cur
}
