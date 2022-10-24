package main

import "container/list"

// LRU缓存算法
func main() {

}

//type ListNodeLRU struct {
//	key, value int
//	prev, next *ListNodeLRU
//}
//
//type LRUCache0201 struct {
//	capacity int
//	head     *ListNodeLRU
//	tail     *ListNodeLRU
//	cache    map[int]*ListNodeLRU
//}
//
//// Constructor01 (int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//func Constructor01(capacity int) LRUCache0201 {
//	head := &ListNodeLRU{}
//	tail := &ListNodeLRU{}
//	head.next = tail
//	tail.prev = head
//	return LRUCache0201{
//		cache:    map[int]*ListNodeLRU{},
//		head:     head,
//		tail:     tail,
//		capacity: capacity,
//	}
//}
//
//// Get 存在返回值,并提升到链表最前 不存在 返回-1
//func (lru *LRUCache0201) Get(key int) int {
//	if res, ok := lru.cache[key]; ok { // exist
//		lru.moveToHead(res)
//		return res.value
//	}
//	return -1
//}
//
//// Put 存在则改值,并提升到链头,不存在,看数量是否够用,不够用再删除再加入
//func (lru *LRUCache0201) Put(key int, value int) {
//	if res, ok := lru.cache[key]; ok { // exist
//		lru.cache[key].value = value
//		lru.moveToHead(res)
//		return
//	}
//	newNode := &ListNodeLRU{
//		key:   key,
//		value: value,
//	}
//	if lru.capacity == len(lru.cache) {
//		delete(lru.cache, lru.removeTail())
//	}
//	lru.addToHead(newNode)
//	lru.cache[key] = newNode
//}
//
//func (lru *LRUCache0201) moveToHead(res *ListNodeLRU) {
//	lru.deleteNode(res)
//	lru.addToHead(res)
//}
//
//func (lru *LRUCache0201) deleteNode(res *ListNodeLRU) {
//	res.prev.next = res.next
//	res.next.prev = res.prev
//}
//
//func (lru *LRUCache0201) addToHead(res *ListNodeLRU) {
//	res.next = lru.head.next
//	res.prev = lru.head
//	lru.head.next.prev = res
//	lru.head.next = res
//}
//
//func (lru *LRUCache0201) removeTail() int {
//	node := lru.tail.prev
//	lru.deleteNode(node)
//	return node.key
//}

// LRUCache02 方法二:go 内部list包
type LRUCache02 struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache02 {
	return LRUCache02{
		capacity,
		map[int]*list.Element{},
		list.New(),
	}
}

type entry struct {
	key, value int
}

func (lru *LRUCache02) Get(key int) int {
	if res, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(res)
		return res.Value.(entry).value
	}
	return -1
}

func (lru *LRUCache02) Put(key int, value int) {
	if res, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(res)
		res.Value = entry{key, value}
		return
	}
	if lru.capacity == len(lru.cache) {
		delete(lru.cache, lru.list.Remove(lru.list.Back()).(entry).key)
	}
	lru.cache[key] = lru.list.PushFront(entry{key, value})
}
