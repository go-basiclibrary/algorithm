package main

// 24 两两反转链表
func main() {
}

/*
	两种方法:
		1.递归
		2.动态规划
*/
type ListNode02 struct {
	Val  int
	Next *ListNode02
}

func swapPairs0201(head *ListNode02) *ListNode02 {
	if head == nil || head.Next == nil { // 剩余一个或者没有了
		return head
	}
	newHead := swapPairs0201(head.Next.Next) //反转后的新链
	prev := head.Next
	prev.Next = head
	head.Next = newHead

	return prev
}

// 动态规划,虚拟出一个新的空节点
func swapPairs0202(head *ListNode) *ListNode {
	prev := &ListNode{}
	prev.Next = head
	endHead := prev
	for prev.Next != nil && prev.Next.Next != nil {
		node1 := prev.Next
		node2 := node1.Next
		subHead := node2.Next

		node2.Next = node1
		node1.Next = subHead
		prev.Next = node2

		//指向下一个开始节点
		prev = node1
	}

	return endHead.Next
}
