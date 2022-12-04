package main

// MyCircularDeque 641 设计循环双端队列
//type MyCircularDeque struct {
//	cap  int
//	list []int
//}
//
//func Constructor01(k int) MyCircularDeque { //构造函数,双端队列最大为 k
//	return MyCircularDeque{
//		k,
//		make([]int, 0),
//	}
//}
//
//// InsertFront 将一个元素添加到双端队列头部.如果操作成功返回 true,否则返回false
//func (md *MyCircularDeque) InsertFront(value int) bool {
//	if len(md.list) == md.cap {
//		return false
//	}
//	ints := make([]int, 1, len(md.list)+1)
//	ints[0] = value
//	md.list = append(ints, md.list...)
//	return true
//}
//
//// InsertLast 将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false
//func (md *MyCircularDeque) InsertLast(value int) bool {
//	if len(md.list) == md.cap {
//		return false
//	}
//	md.list = append(md.list, value)
//	return true
//}
//
//// DeleteFront 从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false
//func (md *MyCircularDeque) DeleteFront() bool {
//	if md.IsEmpty() {
//		return false
//	}
//	md.list = md.list[1:]
//	return true
//}
//
//// DeleteLast 从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false
//func (md *MyCircularDeque) DeleteLast() bool {
//	if md.IsEmpty() {
//		return false
//	}
//	md.list = md.list[:len(md.list)-1]
//	return true
//}
//
//// GetFront 从双端队列头部获得一个元素。如果双端队列为空，返回 -1
//func (md *MyCircularDeque) GetFront() int {
//	if md.IsEmpty() {
//		return -1
//	}
//	return md.list[0]
//}
//
//// GetRear 获得双端队列的最后一个元素。 如果双端队列为空，返回 -1
//func (md *MyCircularDeque) GetRear() int {
//	if md.IsEmpty() {
//		return -1
//	}
//	return md.list[len(md.list)-1]
//}
//
//// IsEmpty 若双端队列为空，则返回 true ，否则返回 false
//func (md *MyCircularDeque) IsEmpty() bool {
//	return len(md.list) == 0
//}
//
//// IsFull 若双端队列满了，则返回 true ，否则返回 false
//func (md *MyCircularDeque) IsFull() bool {
//	return md.cap == len(md.list)
//}

type Node struct {
	val        int
	prev, next *Node
}

// MyCircularDeque 链表
type MyCircularDeque struct {
	head, tail *Node
	size       int
	curSize    int
}

func Constructor01(k int) MyCircularDeque {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return MyCircularDeque{
		head, tail,
		k, 0,
	}
}

func (md *MyCircularDeque) InsertFront(value int) bool {
	if md.size == md.curSize {
		return false
	}
	newN := &Node{val: value}
	newN.next = md.head.next
	newN.prev = md.head
	md.head.next.prev = newN
	md.head.next = newN
	md.curSize++
	return true
}

func (md *MyCircularDeque) InsertLast(value int) bool {
	if md.size == md.curSize {
		return false
	}
	newN := &Node{val: value}
	newN.next = md.tail
	newN.prev = md.tail.prev
	md.tail.prev.next = newN
	md.tail.prev = newN
	md.curSize++
	return true
}

func (md *MyCircularDeque) DeleteFront() bool {
	if md.curSize == 0 {
		return false
	}
	md.head.next = md.head.next.next
	md.head.next.prev = md.head
	md.curSize--
	return true
}

func (md *MyCircularDeque) DeleteLast() bool {
	if md.curSize == 0 {
		return false
	}
	md.tail.prev = md.tail.prev.prev
	md.tail.prev.next = md.tail
	md.curSize--
	return true
}

func (md *MyCircularDeque) GetFront() int {
	if md.IsEmpty() {
		return -1
	}
	return md.head.next.val
}

func (md *MyCircularDeque) GetRear() int {
	if md.IsEmpty() {
		return -1
	}
	return md.tail.prev.val
}

func (md *MyCircularDeque) IsEmpty() bool {
	return md.curSize == 0
}

func (md *MyCircularDeque) IsFull() bool {
	return md.curSize == md.size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
