package Medium

func subsets(nums []int) [][]int {
	var res [][]int
	var subset []int

	var backtrack func(idx int)

	backtrack = func(idx int) {
		if idx >= len(nums) {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			res = append(res, subsetCopy)
			return
		}

		subset = append(subset, nums[idx])
		backtrack(idx + 1)

		subset = subset[:len(subset)-1]
		backtrack(idx + 1)
	}

	backtrack(0)

	return res
}
