package main

// MyCircularDeque 641 设计循环双端队列
// 链表方式实现
type MyCircularDeque struct { // 双端队列最大为k
	head, tail     *Node // 头尾节点
	capacity, size int   // 容量,当前容量
}

type Node struct {
	prev, next *Node
	val        int
}

func ConstructorMCD(k int) MyCircularDeque {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return MyCircularDeque{
		head, tail, k, 0,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool { // 将一个元素添加到双端队列头部
	// success true failed false
	if d.size == d.capacity {
		return false
	}
	newNode := &Node{val: value}
	// 接入
	newNode.prev = d.head
	newNode.next = d.head.next
	d.head.next.prev = newNode
	d.head.next = newNode

	// size ++
	d.size++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool { // 将一个元素添加到双端队列尾部
	// success true failed false
	if d.size == d.capacity {
		return false
	}

	newNode := &Node{val: value}
	// 接入
	newNode.prev = d.tail.prev
	newNode.next = d.tail
	d.tail.prev.next = newNode
	d.tail.prev = newNode

	// size ++
	d.size++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool { //从双端队列头部删除一个元素
	// success true failed false
	if d.size == 0 {
		return false
	}
	// 接出
	delNode := d.head.next
	d.head.next = d.head.next.next
	d.head.next.prev = d.head
	delNode.prev = nil
	delNode.next = nil

	// size --
	d.size--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool { //从双端队列尾部删除一个元素
	// success true failed false
	if d.size == 0 {
		return false
	}
	// 接出
	delNode := d.tail.prev
	d.tail.prev = d.tail.prev.prev
	d.tail.prev.next = d.tail
	delNode.prev = nil
	delNode.next = nil

	// size --
	d.size--
	return true
}

func (d *MyCircularDeque) GetFront() int { // 从双端队列头部获取一个元素
	// val or -1
	if d.size == 0 {
		return -1
	}

	return d.head.next.val
}

func (d *MyCircularDeque) GetRear() int { // 获取双端队列最后一个元素
	// var or -1
	// val or -1
	if d.size == 0 {
		return -1
	}

	return d.tail.prev.val
}

func (d *MyCircularDeque) IsEmpty() bool { // 双端队列为空
	// true or false
	if d.size == 0 {
		return true
	}
	return false
}

func (d *MyCircularDeque) IsFull() bool { // 双端队列满了
	// ture or false
	if d.size == d.capacity {
		return true
	}
	return false
}
