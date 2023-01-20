package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

// N皇后
// 如何判定在同一斜线上,相差横纵相等(左斜线)
// 右斜线,行下标与列下标之和相等
var solutions [][]string //最终答案

func solveNQueens(n int) [][]string {
	//初始化solutions
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1 //先给皇后初始化位置-1
	}
	columns := map[int]bool{}                                // 记录每一列是否有皇后
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{} //两个对角线
	backtrackQ(queens, n, 0, columns, diagonals1, diagonals2)

	return solutions
}

func backtrackQ(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n) // 初始化board
		solutions = append(solutions, board)
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i // 左斜线求差  判定是否在同一斜线
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i // 右斜线求和  判定是否在同一斜线
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i   // 该行放一个皇后
		columns[i] = true // 记录该列已存在皇后
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		// 进入下一行寻找皇后
		backtrackQ(queens, n, row+1, columns, diagonals1, diagonals2)
		// 这里对于回溯要做的操作,删除该节点的皇后需要
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
