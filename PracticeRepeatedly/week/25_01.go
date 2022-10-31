package main

func main() {
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	storeHead := head
	for k > count {
		if head == nil {
			return storeHead
		}
		head = head.Next
		count++
	}

	newList := reverseKGroup(head, k)
	localReverse := reverseListNow(storeHead, k)
	storeHead.Next = newList
	return localReverse
}

func reverseListNow(head *ListNode, k int) *ListNode {
	var newList *ListNode
	prev := head
	for i := 0; i < k; i++ {
		nextNode := prev.Next
		prev.Next = newList
		newList = prev
		prev = nextNode
	}

	return newList
}
