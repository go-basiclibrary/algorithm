package main

import "fmt"

func main() {
	fmt.Println(solveNQueens02(4))
}

var resQ [][]string

// N皇后
// 如何判定在同一斜线上,相差横纵相等(左斜线)
// 右斜线,行下标与列下标之和相等
func solveNQueens02(n int) [][]string {
	// 初始化每一个棋盘
	resQ = [][]string{}
	// 初始化皇后
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}                                // 记录当前列的皇后
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{} // 记录对角线
	backtrackQueens(queens, n, 0, columns, diagonals1, diagonals2)

	return resQ
}

func backtrackQueens(queens []int, n, row int, columns map[int]bool, diagonals1, diagonals2 map[int]bool) {
	// 结束条件
	if row == n {
		//解析答案
		resQ = append(resQ, generateBoardQ(queens, n))
	}
	// 每行递归,找位置
	for i := 0; i < n; i++ {
		// 如果当前列存在,则到下一个格子
		if columns[i] {
			continue
		}
		diagonal1 := row - i // 判断左斜线
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i // 判断右斜线
		if diagonals2[diagonal2] {
			continue
		}
		// 如果上述结果都通过,那么该位置是合理存在的
		columns[i] = true // 该列位置放上皇后
		queens[row] = i   // 皇后的位置
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		// 进入下一层,寻找下一层的皇后
		backtrackQueens(queens, n, row+1, columns, diagonals1, diagonals2)
		// 回溯针对每一层数据进行回退
		delete(columns, i)
		queens[row] = -1
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

// 解析每一个棋盘的结果
func generateBoardQ(queens []int, n int) []string {
	board := make([]string, 0, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q' // 第i列的Q
		board = append(board, string(row))
	}
	return board
}
