package main

import "sort"

// 哈希解法
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var max, key int
	for k, v := range m {
		if v > key {
			key = v
			max = k
		}
	}
	return max
}

// 排序 返回中间解法
func majorityElement02(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[(len(nums)-1)/2]
}

// map 获取长度
func majorityElement03(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var max int
	for k, v := range m {
		if v > len(nums)/2 {
			max = k
		}
	}
	return max
}

//“同归于尽消杀法” :
//由于多数超过50%, 比如100个数，那么多数至少51个，剩下少数是49个。
//遍历数组
//第一个到来的士兵，直接插上自己阵营的旗帜占领这块高地，此时领主 winner 就是这个阵营的人，现存兵力 count = 1。
//如果新来的士兵和前一个士兵是同一阵营，则集合起来占领高地，领主不变，winner 依然是当前这个士兵所属阵营，现存兵力 count 加一；
//如果新来到的士兵不是同一阵营，则前方阵营派一个士兵和它同归于尽。 此时前方阵营兵力-1, 即使双方都死光，这块高地的旗帜 winner 不变，没有可以去换上自己的新旗帜。
//当下一个士兵到来，发现前方阵营已经没有兵力，新士兵就成了领主，winner 变成这个士兵所属阵营的旗帜，现存兵力 count ++。
//就这样各路军阀一直厮杀以一敌一同归于尽的方式下去，直到少数阵营都死光，剩下几个必然属于多数阵营的，winner 是多数阵营。
//（多数阵营 51个，少数阵营只有49个，死剩下的2个就是多数阵营的人）
// Boyer-Moore 投票算法
func majorityElement04(nums []int) int {
	var candidate, count int
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count++
		} else if candidate != v {
			count--
		} else if candidate == v {
			count++
		}
	}
	return candidate
}
