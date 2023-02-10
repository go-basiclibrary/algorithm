package main

// 529 扫雷游戏
func updateBoard(board [][]byte, click []int) [][]byte {
	return engine(board, click[0], click[1])
}

// 八个方向坐标
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

// 整个扫雷游戏的引擎
func engine(board [][]byte, x int, y int) [][]byte {
	// 如果一个地雷被挖出,改为X并结束游戏
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(board, x, y)
	}
	return board
}

func dfs(board [][]byte, x, y int) {
	var cnt int
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			cnt++
		}
	}
	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			dfs(board, tx, ty)
		}
	}
}
