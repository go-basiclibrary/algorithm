package main

import "fmt"

func main() {
	res := minDistance("distance", "springbok")
	fmt.Println(res)
}

// 72 编辑距离
func minDistance(word1 string, word2 string) int {
	// 空串处理
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}

	// 初始化DP切片
	dpSlice := make([][]int, len(word1)+1)

	// 初始化dpSlice  以及边界状态
	for i := range dpSlice {
		dpSlice[i] = make([]int, len(word2)+1)
		dpSlice[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dpSlice[0][i] = i
	}

	// 计算所有的dp值
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			up := dpSlice[i-1][j] + 1
			left := dpSlice[i][j-1] + 1
			upLeft := dpSlice[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				upLeft += 1
			}
			dpSlice[i][j] = min1(up, min1(left, upLeft))
		}
	}

	return dpSlice[len(word1)][len(word2)]
}

func min1(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
