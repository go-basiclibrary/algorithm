package main

import "fmt"

// 146 LRU缓存
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Get(1)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

/*
	思路一:map用于缓存key value
		双向链表,解决访问次数问题
*/

type LinkedList struct {
	key   int
	value int
	prev  *LinkedList
	next  *LinkedList
}

type LRUCache struct {
	// 使用map来做缓存
	Cache map[int]*LinkedList
	// cap控制
	CapCity int
	// 计数
	count int

	head *LinkedList
	tail *LinkedList
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Cache:   make(map[int]*LinkedList),
		CapCity: capacity,
		head:    &LinkedList{},
		tail:    &LinkedList{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lru *LRUCache) Get(key int) int {
	//从缓存中拿数据,存在则返回
	if res, ok := lru.Cache[key]; !ok { //不存在
		//不存在返回-1
		return -1
	} else {
		//将该链写入链尾
		lru.popToTail(res)
		return res.value
	}
}

func (lru *LRUCache) Put(key int, value int) {
	entry := &LinkedList{key: key, value: value}
	//存在则更新值
	if _, ok := lru.Cache[key]; ok { // exist
		//改值
		lru.Cache[key] = entry
		//加入表尾
		lru.popToTail(entry)
	} else {
		if lru.CapCity < lru.count+1 {
			//移除表头
			lru.removeHead()
			delete(lru.Cache, key)
			//加入数据到表尾
			lru.popToTail(entry)
		} else {
			//加入数据到表尾
			lru.popToTail(entry)
		}
		lru.Cache[key] = entry
		lru.count++
	}
}

// 将数据搞入链尾
func (lru *LRUCache) popToTail(entry *LinkedList) {
	entry.next = lru.tail
	entry.prev = lru.tail.prev
	lru.tail.prev.next = entry
	lru.tail.prev = entry
}

//  删除表头
func (lru *LRUCache) removeHead() {
	lru.count--
	node := lru.head.next
	lru.deleteNode(node)
}

func (lru *LRUCache) deleteNode(node *LinkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
