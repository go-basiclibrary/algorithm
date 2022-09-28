package main

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func swapPairs(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairs2(head *ListNode1) *ListNode1 {
	dummyHead := &ListNode1{}
	dummyHead.Next = head
	tem := dummyHead
	for tem.Next != nil && tem.Next.Next != nil {
		node1 := tem.Next
		node2 := tem.Next.Next
		tem.Next = node2
		node1.Next = node2.Next
		tem = node1
	}

	return dummyHead.Next
}
