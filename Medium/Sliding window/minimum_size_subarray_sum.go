package Sliding_window

func minSubArrayLen(target int, nums []int) int {
	minSeq := len(nums)
	l, curSum := 0, 0

	for r := 0; r < len(nums); r++ {
		curSum += nums[r]

		for curSum >= target {

			if minSeq > r-l+1 {
				minSeq = r - l + 1
			}

			curSum -= nums[l]
			l++
		}
	}

	if minSeq == len(nums) && curSum < target {
		return 0
	}

	return minSeq
}
