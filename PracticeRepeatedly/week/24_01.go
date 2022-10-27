package main

// 24  两两交换链表中的节点
// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := swapPairs(head.Next.Next)
	res := head.Next
	head.Next.Next = head
	head.Next = newList

	return res
}

// 三 指针
func swapPairs02(head *ListNode) *ListNode {
	dum := &ListNode{}
	dum.Next = head
	newDum := dum
	for dum.Next != nil && dum.Next.Next != nil {
		node1 := dum.Next
		node2 := dum.Next.Next

		dum.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		dum = node1
	}

	return newDum.Next
}
