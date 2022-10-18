package main

import "container/list"

// LRU缓存算法
func main() {
	c := Constructor(2)
	c.Get(2)
	c.Put(2, 6)
	c.Get(1)
	c.Put(1, 5)
	c.Put(1, 2)
	c.Get(1)
	c.Get(2)
	c.Put(5, 5)
	c.Get(5)
}

/*
	方法一:
		map+双向链表,自实现
*/
//type linkList struct {
//	key, value int
//	next       *linkList
//	prev       *linkList
//}
//
//type LRUCache struct {
//	cache    map[int]*linkList
//	capacity int
//	head     *linkList
//	tail     *linkList
//}
//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		make(map[int]*linkList),
//		capacity,
//		&linkList{},
//		&linkList{},
//	}
//	//头尾行程双向链表结构
//	l.head.next = l.tail
//	l.tail.prev = l.head
//
//	return l
//}
//
//func (l *LRUCache) Get(key int) int {
//	//获取数据,并将其计入链尾,不存在则-1
//	if res, ok := l.cache[key]; ok {
//		//加入链尾
//		l.popToTail(res)
//		return res.value
//	}
//	return -1
//}
//
//func (l *LRUCache) Put(key int, value int) {
//	//加入数据,容量超过,则去除头部数据,加入新数据到链尾
//	//容量够则直接加入链尾
//	if res, ok := l.cache[key]; ok { //存在
//		//放入链尾
//		l.popToTail(res)
//		//修改值map
//		res.value = value
//		return
//	}
//	if l.capacity == len(l.cache) {
//		//去除链头数据,再将新数据加入链尾
//		delKey := l.removeTohead()
//		//del map
//		delete(l.cache, delKey)
//	}
//	newNode := &linkList{key: key, value: value}
//	l.addToTail(newNode)
//	//入map
//	l.cache[key] = newNode
//}
//
//func (l *LRUCache) popToTail(node *linkList) {
//	//计入链尾两步走
//	//1.将其两端节点相接
//	l.deleteNode(node)
//	//2.计入链尾
//	l.addToTail(node)
//}
//
//func (l *LRUCache) deleteNode(node *linkList) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (l *LRUCache) addToTail(node *linkList) {
//	node.next = l.tail
//	node.prev = l.tail.prev
//	l.tail.prev.next = node
//	l.tail.prev = node
//}
//
//// 移除头部元素
//func (l *LRUCache) removeTohead() int {
//	//不清楚node的生命周期什么时候结束  要么就引入一个context
//	node := l.head.next
//	l.deleteNode(node)
//
//	return node.key
//}

/*
	方法二:map+list
*/
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		map[int]*list.Element{},
		list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(entry).value
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if elem, ok := l.cache[key]; ok {
		elem.Value = entry{key, value}
		l.list.MoveToFront(elem)
		return
	}
	l.cache[key] = l.list.PushFront(entry{key, value})
	if l.capacity < l.list.Len() {
		delete(l.cache, l.list.Remove(l.list.Back()).(entry).key)
	}
}
