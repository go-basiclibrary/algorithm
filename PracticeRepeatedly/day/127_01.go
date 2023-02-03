package main

import "math"

// 127  单词接龙
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(word string) int {
		w, ok := wordId[word]
		if !ok {
			// save word
			w = len(wordId)
			wordId[word] = w
			graph = append(graph, []int{})
		}
		return w
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*' // *it h*t hi*
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2) // 将该单词,连接已存在的 *边,或者存入的 *边
			graph[id2] = append(graph[id2], id1) // 将 *边与该单词进行连接
			s[i] = b                             // 恢复原样
		}
		return id1
	}
	for _, v := range wordList {
		addEdge(v)
	}
	beginId := addEdge(beginWord)
	// 如果endWord不存在,则退出
	endId, ok := wordId[endWord]
	if !ok {
		return 0
	}

	const inf = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0 // 起始向下位置
	queue := []int{beginId}
	for len(queue) > 0 { // BFS 算法,查找最短路径
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}
