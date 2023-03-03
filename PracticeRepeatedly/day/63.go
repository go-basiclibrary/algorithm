package main

// 63 动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var res = make([]int, len(obstacleGrid[0]))
	if obstacleGrid[0][0] == 0 {
		res[0] = 1
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				res[j] = 0
				continue
			}

			if j-1 >= 0 && obstacleGrid[i][j] == 0 {
				res[j] += res[j-1]
			}
		}
	}
	return res[len(res)-1]
}
