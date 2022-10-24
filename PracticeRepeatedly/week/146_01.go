package main

import "container/list"

// LRUCache 146 LRU缓存算法

// 常规双向链表解法
//type LRUCache struct {
//	capacity   int
//	cache      map[int]*doubleLinkedList
//	head, tail *doubleLinkedList
//}
//
//type doubleLinkedList struct {
//	key, value int
//	prev, next *doubleLinkedList
//}
//
//func Constructor(capacity int) LRUCache {
//	head := &doubleLinkedList{}
//	tail := &doubleLinkedList{}
//	head.next = tail
//	tail.prev = head
//	return LRUCache{
//		head:     head,
//		tail:     tail,
//		capacity: capacity,
//		cache:    map[int]*doubleLinkedList{},
//	}
//}
//
//func (lru *LRUCache) Get(key int) int {
//	//存在,数据加入表头,返回该值
//	if res, ok := lru.cache[key]; ok {
//		lru.addToHead(res)
//		return res.value
//	}
//	//不存在返回-1
//	return -1
//}
//
//func (lru *LRUCache) Put(key int, value int) {
//	//存在,加入表头,修改该值
//	if res, ok := lru.cache[key]; ok {
//		res.value = value
//		lru.addToHead(res)
//		return
//	}
//	//不存在,查看链表是否等于容量,等于删除链
//	if lru.capacity == len(lru.cache) {
//		delete(lru.cache, lru.removeTail())
//	}
//	//加入数据到链,接入map
//	newNode := &doubleLinkedList{key: key, value: value}
//	lru.addHead(newNode)
//	lru.cache[key] = newNode
//}
//
//func (lru *LRUCache) addToHead(res *doubleLinkedList) {
//	lru.deleteNode(res)
//	lru.addHead(res)
//}
//
//func (lru *LRUCache) removeTail() int {
//	node := lru.tail.prev
//	lru.deleteNode(node)
//	return node.key
//}
//
//func (lru *LRUCache) deleteNode(res *doubleLinkedList) {
//	res.prev.next = res.next
//	res.next.prev = res.prev
//}
//
//func (lru *LRUCache) addHead(res *doubleLinkedList) {
//	res.prev = lru.head
//	res.next = lru.head.next
//	lru.head.next.prev = res
//	lru.head.next = res
//}

// LRUCache 接入go list,实现更加简单
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		map[int]*list.Element{},
		list.New(),
	}
}

type entry struct {
	key, value int
}

func (lru *LRUCache) Get(key int) int {
	if e, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(e)
		return e.Value.(entry).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if e, ok := lru.cache[key]; ok {
		e.Value = entry{key, value}
		lru.list.MoveToFront(e)
		return
	}
	if lru.capacity == len(lru.cache) {
		delete(lru.cache, lru.list.Remove(lru.list.Back()).(entry).key)
	}
	lru.cache[key] = lru.list.PushFront(entry{key, value})
}
