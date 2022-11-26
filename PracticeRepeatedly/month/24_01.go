package main

func swapPairs(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}
	newList := swapPairs(head.Next.Next)
	retNode := head.Next
	//做反转动作
	//接回后面被修改的链
	head.Next.Next = head
	head.Next = newList
	return retNode
}

// 动态规划
func swapPairs02(head *ListNode) *ListNode {
	prev := &ListNode{}
	prev.Next = head
	newList := prev
	for prev.Next != nil && prev.Next.Next != nil { //保证有两个元素
		n1 := prev.Next
		n2 := n1.Next

		// 交换 n2->n1 n1->下一次的起点 prev.Next=n2连接后续的起点 2->1->4
		n1.Next = n2.Next
		n2.Next = n1

		prev.Next = n2
		prev = n1
	}

	return newList.Next
}
