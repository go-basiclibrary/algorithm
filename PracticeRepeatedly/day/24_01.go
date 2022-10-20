package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

// 24 两两交换链表中的节点
func main() {
	stack := arraystack.New()
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	if pop, ok := stack.Pop(); ok {
		fmt.Println(pop)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	递归
*/ //1->2->3->4
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	subRes := swapPairs(head.Next.Next) //反转 3->4
	headNext := head.Next               //存储 x y new  y
	headNext.Next = head
	head.Next = subRes

	return headNext
}

// 迭代  虚拟前指针
func swapPairs02(head *ListNode) *ListNode {
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

		//prev指向下一个进攻的节点
		prev = node1
	}

	return endHead.Next
}
