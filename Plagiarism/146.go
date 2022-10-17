package main

// LRUCache LRU缓存算法
type LRUCache struct {
	capacity   int
	m          map[int]*Node
	head, tail *Node
}

type Node struct {
	Key       int
	Value     int
	Pre, Next *Node
}

func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.m[key]; ok {
		lru.moveToHead(v)
		return v.Value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if v, ok := lru.m[key]; ok {
		v.Value = value
		lru.moveToHead(v)
		return
	}

	if lru.capacity == len(lru.m) {
		delKey := lru.removeTail()
		delete(lru.m, delKey)
	}

	newNode := &Node{Key: key, Value: value}
	lru.addToHead(newNode)
	lru.m[key] = newNode
}

func (lru *LRUCache) moveToHead(v *Node) {
	lru.deleteNode(v)
	lru.addToHead(v)
}

func (lru *LRUCache) deleteNode(v *Node) {
	v.Pre.Next = v.Next
	v.Next.Pre = v.Pre
}

func (lru *LRUCache) addToHead(v *Node) {
	v.Next = lru.head.Next
	v.Pre = lru.head
	lru.head.Next.Pre = v
	lru.head.Next = v
}

func (lru *LRUCache) removeTail() int {
	node := lru.tail.Pre
	lru.deleteNode(node)
	return node.Key
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}
