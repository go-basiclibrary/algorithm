package main

// 21  合并两个有序链表
// 常规迭代  O(n)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	res := newList

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			newList.Next = list2
			list2 = list2.Next
		} else {
			newList.Next = list1
			list1 = list1.Next
		}
		newList = newList.Next
	}
	if list1 != nil {
		newList.Next = list1
	}
	if list2 != nil {
		newList.Next = list2
	}

	return res.Next
}

// 递归  O(n)
func mergeTwoLists02(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}
