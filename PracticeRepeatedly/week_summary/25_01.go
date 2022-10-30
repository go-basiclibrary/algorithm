package main

// 25 K个一组翻转链表
// 递归,最小子问题
func reverseKGroup(head *ListNode, k int) *ListNode {
	oldList := head
	count := 0
	for count < k {
		if head != nil {
			head = head.Next
		} else {
			return oldList // 原链返回
		}
		count++
	}

	// 被变更前的头指向尾节点
	newList := reverseKGroup(head, k)
	// 反转当前结点链表
	jubuF := reverseList01(oldList, k)

	oldList.Next = newList

	return jubuF
}

func reverseList01(list *ListNode, k int) *ListNode {
	var newList *ListNode
	cur := list
	for i := 0; i < k; i++ {
		node := cur.Next
		cur.Next = newList
		newList = cur
		cur = node
	}

	return newList
}
