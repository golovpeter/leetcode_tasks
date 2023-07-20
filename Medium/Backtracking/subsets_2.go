package Backtracking

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var backtrack func(idx int, subset []int)

	sort.Ints(nums)

	backtrack = func(idx int, subset []int) {
		if idx >= len(nums) {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			res = append(res, subsetCopy)
			return
		}

		subset = append(subset, nums[idx])
		backtrack(idx+1, subset)

		subset = subset[:len(subset)-1]

		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx++
		}

		backtrack(idx+1, subset)
	}

	backtrack(0, make([]int, 0))

	return res
}
