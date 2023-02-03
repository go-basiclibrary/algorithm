package main

import "fmt"

func main() {
	islands := numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	})
	fmt.Println(islands)
}

// 200 岛屿数量
// DFS
func numIslands(grid [][]byte) int {
	var variation func(x, y, xLen, yLen int)
	variation = func(x, y, xLen, yLen int) {
		// terminate
		if x < 0 || y < 0 || x == xLen || y == yLen || grid[x][y] == '0' {
			return
		}
		// current logic
		if grid[x][y] == '1' {
			// 感染变异
			grid[x][y] = '0'
		}
		// drill down
		variation(x+1, y, xLen, yLen) // 向下扩展
		variation(x, y+1, xLen, yLen) // 向右扩展
		variation(x-1, y, xLen, yLen) // 向上扩展
		variation(x, y-1, xLen, yLen) // 向左扩展
		// clear current
	}

	var isLandNums int
	for x, gr := range grid {
		for y, v := range gr {
			if v == '1' {
				// 这里我们要做一个事情,那就是将周围所有挨着的岛屿变为0
				variation(x, y, len(grid), len(gr))
				isLandNums++
			}
		}
	}
	return isLandNums
}
