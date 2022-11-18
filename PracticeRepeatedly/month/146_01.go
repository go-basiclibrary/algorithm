package main

import "container/list"

// LRUCache
// 146 LRU缓存算法
// 先去缓存(高一级的存储器都可以称作低一级存储器的缓存)中查找,不存在,则去内存查找,找到,存入缓存,返回
// map + 链表
//type LRUCache struct {
//	cache      map[int]*List
//	cap        int
//	head, tail *List
//}
//
//type List struct {
//	prev, next *List
//	key, val   int
//}
//
//func Constructor(capacity int) LRUCache { // capacity 初始容量
//	head := &List{}
//	tail := &List{}
//	head.next = tail
//	tail.prev = head
//	return LRUCache{
//		make(map[int]*List),
//		capacity,
//		head, tail,
//	}
//}
//
//func (l *LRUCache) Get(key int) int { // 返回值,并且设置为最近使用,不存在 return-1
//	if res, ok := l.cache[key]; ok {
//		// 存在,设置为常用值
//		// 断链,接链
//		l.delNode(res)
//		l.accessNode(res)
//		return res.val
//	}
//	return -1
//}
//
//func (l *LRUCache) accessNode(res *List) {
//	//放入链尾为最常用元素
//	res.next = l.tail
//	res.prev = l.tail.prev
//	l.tail.prev.next = res
//	l.tail.prev = res
//}
//
//func (l *LRUCache) delNode(res *List) int {
//	res.prev.next = res.next
//	res.next.prev = res.prev
//	res.next = nil
//	res.prev = nil
//	return res.key
//}
//
//func (l *LRUCache) Put(key int, value int) { // key存在,变更value,不存在,写入,超出capacity,移除最久未使用
//	if res, ok := l.cache[key]; ok {
//		// 变更value
//		res.val = value
//		// 设置为常用值
//		l.delNode(res)
//		l.accessNode(res)
//		return
//	}
//	//不存在
//	// 容量达到阈值 del 不常用值
//	if len(l.cache) == l.cap {
//		delete(l.cache, l.delNode(l.head.next))
//	}
//	// 新建值写入,及设置常用值
//	newNode := &List{key: key, val: value}
//	l.cache[key] = newNode
//	l.accessNode(newNode)
//}

// LRUCache 接入go list
type LRUCache struct {
	list  *list.List
	cache map[int]*list.Element
	cap   int
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list.New(),
		make(map[int]*list.Element),
		capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if res, ok := l.cache[key]; ok {
		l.list.MoveToFront(res)
		return res.Value.(entry).value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if res, ok := l.cache[key]; ok {
		l.list.MoveToFront(res)
		res.Value = entry{key, value}
		return
	}
	if len(l.cache) == l.cap {
		delete(l.cache, l.list.Remove(l.list.Back()).(entry).key)
	}
	l.cache[key] = l.list.PushFront(entry{key, value})
}
