package main

func maximalSquare(matrix [][]byte) int {
	var maxSide int

	dp := make([][]int, len(matrix))

	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if dp[i][j] == 1 { // 等于1 才能算作右下角
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				maxSide = max(dp[i][j], maxSide)
			}
		}
	}

	return maxSide * maxSide
}
