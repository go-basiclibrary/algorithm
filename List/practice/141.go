package main

import "fmt"

//141 环形链表  使用快慢指针法  算法时间复杂度O(n)
func main() {
	fmt.Println(&ListNode141{} == nil)
}

type ListNode141 struct {
	Val  int
	Next *ListNode141
}

// 性能较差
func hasCycle(head *ListNode141) bool {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
