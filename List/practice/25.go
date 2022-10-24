package main

//25 K个一组翻转链表
type ListNode25 struct {
	Val  int
	Next *ListNode25
}

/*
	思路一:
		过滤所有k个节点,有K则进一步进入递归,无k个就结束返回
		翻转所有的内部的K个节点,回归到链表子问题,链表反转
		出去之后,要拿旧链接入新链
*/
// 时间复杂度O(n)
func reverseKGroup(head *ListNode25, k int) *ListNode25 {
	newHead := head
	count := 0
	for count < k {
		if newHead == nil {
			return head
		}
		count++
		newHead = newHead.Next
	}

	newList := reverseKGroup(newHead, k) //这里获取的是反转后的新list
	//翻转
	reverList := reverseList25(head, k)
	//旧list接入新list
	head.Next = newList
	return reverList
}

func reverseList25(head *ListNode25, k int) *ListNode25 {
	var newList *ListNode25
	cur := head
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = newList
		newList = cur
		cur = next
	}

	return newList
}
