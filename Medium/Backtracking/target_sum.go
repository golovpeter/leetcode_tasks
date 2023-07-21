package Backtracking

func findTargetSumWays(nums []int, target int) int {
	var backtrack func(int, int) int

	backtrack = func(currIdx int, target int) int {
		if currIdx > len(nums)-1 {
			if target == 0 {
				return 1
			}

			return 0
		}

		return backtrack(currIdx+1, target-nums[currIdx]) + backtrack(currIdx+1, target+nums[currIdx])
	}

	return backtrack(0, target)
}
