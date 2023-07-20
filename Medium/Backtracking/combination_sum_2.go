package Backtracking

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(start int, currComb []int, currSum int)
	sort.Ints(candidates)

	backtrack = func(start int, currComb []int, currSum int) {
		if currSum > target {
			return
		}

		if currSum == target {
			res = append(res, append([]int{}, currComb...))
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target || (i != start && candidates[i] == candidates[i-1]) {
				continue
			}

			currComb = append(currComb, candidates[i])
			backtrack(i+1, currComb, currSum+candidates[i])
			currComb = currComb[:len(currComb)-1]
		}

		return
	}

	backtrack(0, []int{}, 0)
	return res
}
