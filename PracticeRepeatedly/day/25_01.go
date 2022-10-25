package main

import "fmt"

// ListNode25
// 25 K个一组翻转链表
func main() {
	fmt.Println(reverseKGroup(&ListNode25{Val: 1, Next: &ListNode25{2, &ListNode25{3, &ListNode25{4, &ListNode25{5, nil}}}}}, 2))
}

type ListNode25 struct {
	Val  int
	Next *ListNode25
}

func reverseKGroup(head *ListNode25, k int) *ListNode25 {
	count := 0
	newHead := head
	for count < k {
		if newHead == nil {
			return head
		}
		count++
		newHead = newHead.Next
	}
	newList := reverseKGroup(newHead, k) //反转后的链,被返回
	jubuF := reverseListWithK(head, k)   //这里返回单K链中的翻转  如1->2->3 翻转两个 则返回 2->1

	head.Next = newList
	return jubuF
}

func reverseListWithK(head *ListNode25, k int) *ListNode25 {
	var cur *ListNode25
	prev := head
	for i := 0; i < k; i++ {
		node := prev.Next
		prev.Next = cur
		cur = prev
		prev = node
	}
	return cur
}
