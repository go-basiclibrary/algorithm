package main

// 21 合并两个有序链表
// 迭代
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	newList := res
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			res.Next = list2
			list2 = list2.Next
		} else {
			res.Next = list1
			list1 = list1.Next
		}
		res = res.Next
	}

	if list1 != nil {
		res.Next = list1
	}
	if list2 != nil {
		res.Next = list2
	}
	return newList.Next
}

// 迭代
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
