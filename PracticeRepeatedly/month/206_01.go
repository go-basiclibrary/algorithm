package main

import "fmt"

// ListNode 206 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data := reverseList02(&ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	fmt.Println(data)
}

// 迭代法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}
	return pre
}

// 递归
func reverseList02(head *ListNode) *ListNode {
	//终结条件
	if head == nil || head.Next == nil {
		return head
	}
	//进入到下一层
	node := reverseList02(head.Next)
	//逻辑
	head.Next.Next = head
	//清理当前层
	head.Next = nil //主要考虑的是,最终,1->2的时候,形成环,没有被解决掉
	return node
}

//妖魔化双指针
func reverseList03(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for head.Next != nil {
		// 存储下一个起点
		t := head.Next.Next
		head.Next.Next = cur
		cur = head.Next
		head.Next = t
	}
	return cur
}
