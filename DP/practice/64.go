package main

import "math"

// dfs func  超时
func minPathSum(grid [][]int) int {
	var minPath = math.MaxInt

	var dfs func(grid [][]int, x, y, pathSum int)

	dfs = func(grid [][]int, x, y, pathSum int) {
		// terminate
		if x >= len(grid) || y >= len(grid[0]) {
			return
		}

		if x == len(grid)-1 && y == len(grid[0])-1 {
			// go to end
			pathSum += grid[x][y]
			minPath = min(minPath, pathSum)
			return
		}

		// current logic
		pathSum += grid[x][y]

		// drill down
		dfs(grid, x, y+1, pathSum) // 向左移动
		dfs(grid, x+1, y, pathSum) // 向下移动
	}

	dfs(grid, 0, 0, 0)

	return minPath
}

// dp func
func minPathSum02(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}

	return dp[len(dp)][len(dp[0])]
}
