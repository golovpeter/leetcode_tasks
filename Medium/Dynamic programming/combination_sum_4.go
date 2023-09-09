package Dynamic_programming

func combinationSum4(nums []int, target int) int {
	dp := map[int]int{
		0: 1,
	}

	for i := 1; i < target+1; i++ {
		dp[i] = 0

		for _, el := range nums {
			dp[i] += dp[i-el]
		}
	}

	return dp[target]
}
