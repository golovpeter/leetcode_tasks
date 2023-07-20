package Dynamic_programming

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	result := 0

	for i := 0; i < len(nums); i++ {
		maxSub := 0

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > maxSub {
				maxSub = dp[j]
			}
		}

		dp[i] = maxSub + 1

		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}
