package main

import (
	"fmt"
)

func main() {
	answer := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(answer)
}

// 127 单词接龙
// DFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, v := range wordList {
		addEdge(v)
	}
	beginId := addEdge(beginWord)
	endId, ok := wordId[endWord]
	if !ok {
		return 0
	}

	const inf int = -1
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
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
