package Medium

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		for _, el := range coins {
			if i-el >= 0 {
				dp[i] = min(dp[i], 1+dp[i-el])
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
