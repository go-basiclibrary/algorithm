package main

import "fmt"

// 146 LRU缓存  (map+双向链表)
func main() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(2))
	cache.Put(1, 1)
	cache.Put(4, 1)
	fmt.Println(cache.Get(2))
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
	if res, ok := lru.Cache[key]; ok {
		lru.moveToTail(res)
		return res.value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if res, ok := lru.Cache[key]; ok {
		lru.moveToTail(res)
		lru.Cache[key].value = value
		return
	}

	if lru.CapCity == len(lru.Cache) {
		delKey := lru.removeHead()
		delete(lru.Cache, delKey)
	}

	newNode := &LinkedList{key: key, value: value}
	lru.addToTail(newNode)
	lru.Cache[key] = newNode
}

func (lru *LRUCache) moveToTail(n *LinkedList) {
	//删除节点
	lru.deleteToNode(n)
	//再加入尾
	lru.addToTail(n)
}

func (lru *LRUCache) deleteToNode(n *LinkedList) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *LRUCache) addToTail(n *LinkedList) {
	//接入
	n.next = lru.tail
	n.prev = lru.tail.prev
	//切断
	lru.tail.prev.next = n
	lru.tail.prev = n
}

func (lru *LRUCache) removeHead() int {
	node := lru.head.next
	lru.deleteToNode(node)
	return node.key
}
