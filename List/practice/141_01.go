package main

// 141 环形链表
func main() {
}

type ListNode14101 struct {
	Val  int
	Next *ListNode14101
}

/*
	两种方法
		1.map
		2.快慢指针
*/
func hasCycle01(head *ListNode14101) bool {
	m := make(map[*ListNode14101]int)
	//遍历链表
	for head != nil {
		//判断map是否存在该数据,不存在写入,存在,该链为环状
		// exist
		if _, ok := m[head]; ok {
			return true
		} else {
			m[head] = head.Val
		}
		//变更节点起点
		head = head.Next
	}

	return false
}

// 快慢指针,反人类思维
func hasCycle02(head *ListNode14101) bool {
	quick, slow := head, head
	for quick != nil && slow != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			return true
		}
	}

	return false
}
