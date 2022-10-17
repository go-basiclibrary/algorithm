package main

import "container/list"

// 146 LRU缓存多种解法  (使用map加list)
func main() {

}

type entry struct {
	key, value int
}

type LRUCache02 struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

func Constructor02(capacity int) LRUCache02 {
	return LRUCache02{capacity, map[int]*list.Element{}, list.New()}
}

func (c *LRUCache02) Get(key int) int {
	if res, ok := c.cache[key]; ok {
		c.list.MoveToFront(res)
		return res.Value.(entry).value
	}
	return -1
}

func (c *LRUCache02) Put(key, value int) {
	if res, ok := c.cache[key]; ok {
		res.Value = entry{key, value}
		c.list.MoveToFront(res) //刷新使用时间
		return
	}
	c.cache[key] = c.list.PushFront(entry{key, value})
	if len(c.cache) > c.cap {
		delete(c.cache, c.list.Remove(c.list.Back()).(entry).key)
	}
}
