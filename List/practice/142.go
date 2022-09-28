package main

// 142 环形链表Ⅱ
func main() {

}

// 两个解决方案,时间复杂度都是O(n),空间复杂度都是O(1)
func detectCycle(head *ListNode) *ListNode {
	//双指针都从起点位置开始走
	fast, slow := head, head
	//找到第一次相遇位置然后退出
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	//重置fast,寻求第二次相遇,第二次相遇所走的步数就是环入口
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}

func detectCycle2(head *ListNode) *ListNode {
	//双指针都从起点位置开始走
	fast, slow := head, head
	//找到第一次相遇位置然后退出
	for {
		if !(fast != nil && fast.Next != nil) {
			return nil
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	//重置fast,寻求第二次相遇,第二次相遇所走的步数就是环入口
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}
