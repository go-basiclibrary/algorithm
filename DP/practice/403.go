package main

func canCross(stones []int) bool {
	n := len(stones)
	DP := make([][]bool, n)
	for i := range DP {
		DP[i] = make([]bool, n)
	}
	DP[0][0] = true

	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			DP[i][k] = DP[j][k-1] || DP[j][k] || DP[j][k+1]
			if i == n-1 && DP[i][k] {
				return true
			}
		}
	}
	return false
}
