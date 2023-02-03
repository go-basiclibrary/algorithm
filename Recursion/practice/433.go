package main

import (
	"fmt"
)

func main() {
	fmt.Println(minMutation(
		"AACCGGTT",
		"AACCGGTA",
		[]string{"AACCGGTA"},
	))
}

// 433 最小基因变化 DFS
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	m := make(map[string]struct{})
	for i := 0; i < len(bank); i++ {
		m[bank[i]] = struct{}{}
	}
	if _, ok := m[endGene]; !ok {
		return -1
	}

	q := []string{startGene}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, v := range cur {
				for _, s := range "ACGT" {
					if v != s {
						nxt := cur[:i] + string(s) + cur[i+1:]
						if _, ok := m[nxt]; ok {
							if nxt == endGene {
								return step + 1
							}
							delete(m, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}
