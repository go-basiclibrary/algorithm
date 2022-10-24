package main

// ListNode 24  两两交换
type ListNode struct {
	Val  int
	Next *ListNode
}

//func swapPairs(head *ListNode) *ListNode {
//	dummyHead := &ListNode{}
//	dummyHead.Next = head
//	tem := dummyHead
//	for dummyHead.Next != nil && dummyHead.Next.Next != nil {
//		node1 := dummyHead.Next
//		node2 := dummyHead.Next.Next
//
//		dummyHead.Next = node2
//		node1.Next = node2.Next
//		node2.Next = node1
//
//		dummyHead = node1
//	}
//
//	return tem.Next
//}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	subRes := swapPairs(head.Next.Next)
	headNext := head.Next
	headNext.Next = head
	head.Next = subRes

	return headNext
}
