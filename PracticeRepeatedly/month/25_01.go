package main

// 25 K个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	// 进入下一层
	newList := reverseKGroup(node, k) //反转后的newList
	// 当前层逻辑
	revList := reverse(head, k)
	head.Next = newList
	return revList
}

func reverse(head *ListNode, k int) *ListNode {
	var cur *ListNode
	for i := 0; i < k; i++ {
		node := head.Next
		head.Next = cur
		cur = head
		head = node
	}
	return cur
}
