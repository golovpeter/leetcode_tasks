package Medium

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0

	for _, el := range nums {
		if curSum < 0 {
			curSum = 0
		}

		curSum += el

		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}
