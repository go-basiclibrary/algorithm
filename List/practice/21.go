package main

// ListNode21 21 合并两个有序链表
type ListNode21 struct {
	Val  int
	Next *ListNode21
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//普通迭代
	var newList = &ListNode{}
	cur := newList
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 == nil {
			newList.Next = list1
			list1 = list1.Next
		}
		if list1 == nil && list2 != nil {
			newList.Next = list2
			list2 = list2.Next
		}
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				newList.Next = list2
				list2 = list2.Next
			} else {
				newList.Next = list1
				list1 = list1.Next
			}
		}
		newList = newList.Next
	}
	return cur.Next
}

//优化迭代
func mergeTwoLists02(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList = &ListNode{}
	cur := newList
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
	return cur.Next
}

// 递归
func mergeTwoLists03(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists03(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists03(list1, list2.Next)
		return list2
	}
}
