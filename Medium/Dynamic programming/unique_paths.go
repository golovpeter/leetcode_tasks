package Dynamic_programming

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	if m > 1 {
		dp[1][0] = 1
	}
	if n > 1 {
		dp[0][1] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
