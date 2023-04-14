package main

// 不同路径 Ⅲ  DFS + 回溯
func uniquePathsIII(grid [][]int) int {
	answer := 0
	m, n := len(grid), len(grid[0])

	starti, startj := 0, 0
	empty := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 { // 找打起点
				starti, startj = i, j
				continue
			}
			if grid[i][j] == 0 {
				empty++
			}
		}
	}

	var dfs func(grid [][]int, i, j, step int)
	dfs = func(grid [][]int, i, j, step int) {
		// dfs terminate
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		// 有障碍,不能走
		if grid[i][j] == -1 {
			return
		}
		// 走过不能走
		if grid[i][j] == -2 {
			return
		}

		// 达到终点,不再走
		// 走遍所有空格,是我们想要的路径
		if grid[i][j] == 2 {
			if step == empty {
				answer++
			}
			return
		}

		// 标记走过
		grid[i][j] = -2

		// drill down
		dfs(grid, i+1, j, step+1)
		dfs(grid, i-1, j, step+1)
		dfs(grid, i, j+1, step+1)
		dfs(grid, i, j-1, step+1)

		// clear current
		// 回溯清理,以便于下一个路径使用
		grid[i][j] = 0
	}

	// start dfs
	dfs(grid, starti, startj, 0)

	return answer
}
