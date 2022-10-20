package main

import "fmt"

//206反转链表
func main() {
	res := reverseList4(&ListNode206{Val: 1, Next: &ListNode206{Val: 2, Next: &ListNode206{Val: 3, Next: &ListNode206{Val: 4, Next: &ListNode206{Val: 5}}}}})
	fmt.Println(*res)
}

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

func reverseList(head *ListNode206) *ListNode206 {
	var prev *ListNode206
	cur := head
	for cur != nil {
		//第一个参数交换,后继指针指向前继节点
		//第二个参数交换,前继节点,通过后继指针向前移动
		//第三个参数,考虑一个问题:这里如果我们不引入第三个参数去保存cur.Next节点,那么cur.Next在下一轮迭代中
		//就会变成nil,所以这里引入第三个参数,用于保存后续便利节点
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func reverseList2(head *ListNode206) *ListNode206 {
	var prev *ListNode206
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 递归
func reverseList3(head *ListNode206) *ListNode206 {
	if head == nil || head.Next == nil {
		return head
	}
	var p = reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func reverseList4(head *ListNode206) *ListNode206 {
	var prev *ListNode206
	cur := head
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}

	return prev
}

func reverseList5(head *ListNode206) *ListNode206 {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
